package router

import (
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/controller/api"
)

func InitBaseRouter(Router *gin.Engine) {
	BaseRouter := Router.Group("base")

		//BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)

}

