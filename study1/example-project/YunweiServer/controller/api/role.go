package api

import (
	"fmt"
	"study1/example-project/YunweiServer/model/modelInterface"
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/model/dbModel"
	"study1/example-project/YunweiServer/controller/servers"
)



func GetRoleList(c *gin.Context){
	type searchParams struct {
		dbModel.Authority
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)

	fmt.Printf("%+v \n",sp)
	err, list, total := sp.Authority.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     sp.PageInfo.Page,
			"pageSize": sp.PageInfo.PageSize,
		})

	}

}



//type RegistAndLoginStuct struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//func Login(c *gin.Context) {
//	var L RegistAndLoginStuct
//	_ = c.BindJSON(&L)
//	U := &dbModel.User{Username: L.Username, Password: L.Password}
//	if err, user := U.Login(); err != nil {
//		servers.ReportFormat(c, false, fmt.Sprintf("用户名密码错误或%v", err), gin.H{})
//	} else {
//		tokenNext(c, *user)
//
//	}
//}
