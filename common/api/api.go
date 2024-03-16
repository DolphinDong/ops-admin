package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
)

type Api struct {
	Context *gin.Context
	Logger  *zap.SugaredLogger
	UserKey string
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = GetRequestLogger(c)
	e.UserKey = GetUserKeyFromContext(c)
	return e
}

// GetLogger 获取上下文提供的日志
func (e Api) GetLogger() *zap.SugaredLogger {
	return GetRequestLogger(e.Context)
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			//e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			e.AddError(err)
			break
		}
	}
	return e
}

// 校验数据
func (e *Api) Validate(struc interface{}) *Api {
	err := validator.New().Struct(struc)
	if err != nil {
		msg := "Invalid parameter: "
		es := err.(validator.ValidationErrors)
		errFields := []string{}
		for _, e := range es {
			errFields = append(errFields, fmt.Sprintf("%v(%v)", e.Field(), e.Tag()))
		}
		// 拼接不符合要求的字段
		msg += strings.Join(errFields, ", ")
		e.AddError(errors.New(msg))
	}
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	Error(e.Context, code, err, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	OK(e.Context, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}
