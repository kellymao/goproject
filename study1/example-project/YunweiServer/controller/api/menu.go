package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/middleware"
)
import "study1/example-project/YunweiServer/model/dbModel"
import "study1/example-project/YunweiServer/controller/servers"



func GetMenu(c *gin.Context){

	claims,_ := c.Get("claims")
	claimsuse := claims.(*middleware.CustomClaims)
	menus := &dbModel.Menu{}

	fmt.Println(claimsuse)
	rel ,err := menus.GetMenu(claimsuse.AuthorityId)
	if err!=nil{
		fmt.Println(err)
		servers.ReportFormat(c, false, fmt.Sprintf("获取菜单menu失败%v", err), gin.H{})

	}else {
		servers.ReportFormat(c, true, "获取菜单menu成功", gin.H{"menu": rel })
	}


}
