package api

import (
	"fmt"
	"study1/example-project/YunweiServer/model/modelInterface"
	"github.com/gin-gonic/gin"
	"study1/example-project/YunweiServer/model/dbModel"
	"study1/example-project/YunweiServer/controller/servers"
)



func GetUserList(c *gin.Context){
	type searchParams struct {
		dbModel.User
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)

	fmt.Printf("%+v \n",sp)
	err, list, total := sp.User.GetInfoList(sp.PageInfo)
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

