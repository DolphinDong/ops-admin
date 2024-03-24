package casbin

import (
	"github.com/DolphinDong/ops-admin/common/models/admin"
	"github.com/DolphinDong/ops-admin/pkg/logger"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

var (
	_enforcer *casbin.Enforcer
	once      = sync.Once{}
)

func SetupCasbin(db *gorm.DB) {
	once.Do(func() {
		model, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && ignoreCase(r.act, p.act)
`)
		if err != nil {
			logger.ZapLogger.Fatalf("New casbin model failed: %+v", errors.WithStack(err))
		}
		csbRuleModel := &admin.CasbinRule{}
		adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, csbRuleModel, csbRuleModel.TableName())
		if err != nil {
			logger.ZapLogger.Fatalf("New casbin adapter failed: %+v", errors.WithStack(err))
		}
		enforcer, err := casbin.NewEnforcer(model, adapter)
		if err != nil {
			logger.ZapLogger.Fatalf("New casbin enforcer failed: %+v", errors.WithStack(err))
		}
		enforcer.AddFunction("ignoreCase", IgnoreCaseFunc)
		err = enforcer.LoadPolicy()
		if err != nil {
			logger.ZapLogger.Fatalf("Load casbin policy failed: %+v", errors.WithStack(err))
		}
		_enforcer = enforcer
		go loadPolicyCyclic()
	})
	logger.ZapLogger.Info("Init casbin enforcer success")
}

func GetEnforcer() *casbin.Enforcer {
	return _enforcer
}

func loadPolicyCyclic() {
	for {
		select {
		case <-time.NewTicker(time.Second * 15).C:
			err := _enforcer.LoadPolicy()
			if err != nil {
				logger.ZapLogger.Errorf("Load casbin policy failed: %+v", errors.WithStack(err))
			} else {
				logger.ZapLogger.Info("Load casbin policy success")
			}
		}
	}
}

func IgnoreCase(key1 string, key2 string) bool {
	return strings.ToLower(key1) == strings.ToLower(key2)
}
func IgnoreCaseFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return (bool)(IgnoreCase(name1, name2)), nil
}
