package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "study1/example-project/YunweiServer/model/dbModel"
import "study1/example-project/YunweiServer/controller/servers"



func Get_menutree(c *gin.Context){

	//claims,_ := c.Get("claims")
	//claimsuse := claims.(*middleware.CustomClaims)

	var roleid string
	err := c.ShouldBindJSON(&roleid)
	if err!=nil{
		fmt.Println("get_menutree 解析参数失败！",err)
	}
	menus := &dbModel.Role_menu_tree{}

	rel ,err := menus.Get_role_menutree(roleid)
	if err!=nil{
		fmt.Println(err)
		servers.ReportFormat(c, false, fmt.Sprintf("获取菜单menu失败%v", err), gin.H{})

	}else {
		servers.ReportFormat(c, true, "获取菜单menu成功", gin.H{"menu": rel })
	}


}

func Save_menutree(c *gin.Context){

	//claims,_ := c.Get("claims")
	//claimsuse := claims.(*middleware.CustomClaims)

	var menus  []dbModel.Role_menu_tree

	err := c.ShouldBindJSON(&menus)
	if err!=nil{
		fmt.Println("Save_menutree 解析参数失败！",err)
	}

	m:= &dbModel.Role_to_menu{}
	m.Delete(menus[0].Roleid)  // 删除原权限

	for _,menu  := range menus{

		err = menu.Save()

		if err!=nil{
			break
		}

	}

	if err!=nil{
		fmt.Println(err)
		servers.ReportFormat(c, false, fmt.Sprintf("Save_menutree失败%v", err), gin.H{})

	}else {
		servers.ReportFormat(c, true, "获取菜单menu成功", gin.H{})
	}


}

