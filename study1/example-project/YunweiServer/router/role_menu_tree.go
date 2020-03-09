package router

import (
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/controller/api"
	"study1/example-project/YunweiServer/middleware"
)

func InitMenuTreeRouter(Router *gin.Engine) {
	BaseRouter := Router.Group("role_menu_tree")

	//BaseRouter.POST("regist", api.Regist)
	BaseRouter.Use(middleware.JWTAuth())
	BaseRouter.POST("getmenutree", api.Get_menutree)

}

