package router

import (
	"context"
	"github.com/DolphinDong/ops-admin/common/rpc"
	"github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

func AddOrUpdateApis(engin *gin.Engine) error {
	var req = admin.AddOrUpdateApiReq{}
	routeInfos := engin.Routes()
	for _, routeInfo := range routeInfos {
		s := strings.Split(routeInfo.Handler, ".")
		handler := s[len(s)-1]
		req.Apis = append(req.Apis, &admin.ApiItem{
			Path:    routeInfo.Path,
			Method:  routeInfo.Method,
			Handler: handler,
		})
	}
	_, err := rpc.AdminClient.AddOrUpdateApi(context.Background(), &req)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
