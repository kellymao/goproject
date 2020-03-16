package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "study1/example-project/YunweiServer/model/dbModel"
import "study1/example-project/YunweiServer/controller/servers"



func GetApi_list(c *gin.Context){

	//claims,_ := c.Get("claims")
	//claimsuse := claims.(*middleware.CustomClaims)

	var roleid string
	err := c.ShouldBindJSON(&roleid)
	if err!=nil{
		fmt.Println("get roleid_to_api 解析参数失败！",err)
	}
	menus := &dbModel.Role_to_api{}

	rel ,err := menus.GetApi_list(roleid)
	if err!=nil{
		fmt.Println(err)
		servers.ReportFormat(c, false, fmt.Sprintf("获取菜单roleid_to_api失败%v", err), gin.H{})

	}else {
		servers.ReportFormat(c, true, "获取菜单roleid_to_api成功", gin.H{"menu": rel })
	}


}


func SaveApi_list(c *gin.Context){

	type Checked_role struct{
		Roleid   string   `json:"roleid"`
		Checkids []int `json:"check_ids"`

	}
	var  checked_role   Checked_role
	err := c.ShouldBindJSON(&checked_role)

	if err!=nil{
		fmt.Println("save roleid_to_api 解析参数失败！",err)
	}


	menus := &dbModel.Role_to_api{}

	fmt.Printf("%+v \n",checked_role)
	err = menus.SaveApi_list(checked_role.Roleid,checked_role.Checkids)
	if err!=nil{
		fmt.Println(err)
		servers.ReportFormat(c, false, fmt.Sprintf("保存菜单roleid_to_api失败%v", err), gin.H{})

	}else {
		servers.ReportFormat(c, true, "保存菜单roleid_to_api成功", gin.H{ })
	}




}




