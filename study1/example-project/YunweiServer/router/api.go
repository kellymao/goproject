package router

import (
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/controller/api"
	"study1/example-project/YunweiServer/middleware"
)

// token 过期认证以及api认证 中间件
func InitApiRouter(Router *gin.Engine) {
	BaseRouter := Router.Group("api")

	//BaseRouter.POST("regist", api.Regist)
	BaseRouter.Use(middleware.JWTAuth())
	BaseRouter.POST("getapilist", api.GetApiList)
	BaseRouter.POST("getallapis",api.GetAllapis)

}
