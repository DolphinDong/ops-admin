package middleware

import (
	"github.com/DolphinDong/ops-admin/common/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

func ErrorHandle(c *gin.Context) {
	defer func() {
		r := recover()
		if r != nil {
			var err error
			if e, ok := r.(error); ok {
				err = errors.WithStack(e)
			} else {
				err = errors.Errorf("%v", r)
			}
			logx.Errorf("%+v", err)
			api.Error(c, http.StatusInternalServerError, err, "服务器异常")
		}
	}()
	c.Next()
}
