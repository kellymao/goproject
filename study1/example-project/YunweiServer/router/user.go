package router

import (
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/controller/api"
	"study1/example-project/YunweiServer/middleware"
)

func InitUserRouter(Router *gin.Engine) {
	BaseRouter := Router.Group("user")

	//BaseRouter.POST("regist", api.Regist)
	BaseRouter.Use(middleware.JWTAuth())
	BaseRouter.POST("getuserlist", api.GetUserList)

}


