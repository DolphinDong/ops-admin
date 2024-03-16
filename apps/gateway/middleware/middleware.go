package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(engin *gin.Engine) {
	// 生产request-id 和 logger
	engin.Use(RequestId)
	// 捕获全局错误
	engin.Use(ErrorHandle)
	// 没有路由
	engin.NoRoute(NoRouter)
	// 打印请求日志
	engin.Use(PrintRequestInfo)
	// 跨域处理
	engin.Use(Options)

	// 校验token
	engin.Use(ParseToken)
	// 权限校验
	engin.Use(PermissionCheck)
}
