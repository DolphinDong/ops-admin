package admin

import "github.com/DolphinDong/ops-admin/common/models"

type Api struct {
	models.BseId
	Path    string `gorm:"comment:请求路径;size=255"`
	Method  string `gorm:"comment:请求方法;size=10"`
	Handler string `gorm:"comment:处理函数名称;size:255"`
	models.ModelTime
	models.ControlBy
}

func (Api) TableName() string {
	return "admin_api"
}
