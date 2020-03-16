package dbModel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"study1/example-project/YunweiServer/model/modelInterface"
	"study1/example-project/YunweiServer/init/qmsql"


)

type Api struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Relatemenuid string `json:"relatemenuid"`
}


func (a *Api) GetInfoList(info modelInterface.PageInfo) (err error, apiList []Api, total int) {


	var db *gorm.DB
	if a.Group != ""  ||  a.Path != "" {
		db = qmsql.DEFAULTDB.Model(&Api{}).Where("`path` = ? or `group` = ? " , a.Path,a.Group).Count(&total)
	}else{
		db = qmsql.DEFAULTDB.Model(&Api{}).Count(&total)


	}

	if err = db.Error;err!=nil{
		fmt.Println("获取apilist total失败！",err )
		return
	}

	limit := info.PageSize // 每页展示的条数
	offset := info.PageSize * (info.Page - 1)  //  条数 * （页数page-1） = 当前页page 是从哪条记录开始

	if limit >0 {
		// 客户端获取分页
		err=db.Limit(limit).Offset(offset).Order("id ").Find(&apiList).Error
	}else{
		// 不获取分页
		err=db.Order("id ").Find(&apiList).Error

	}


	if err!=nil{
		fmt.Println("获取apilist 分页失败！",err )
		return
	}

	return





}

func (a *Api) GetAllapis() (err error, apiList []Api){
	err = qmsql.DEFAULTDB.Find(&apiList).Error
	if err!=nil{
		fmt.Println("获取allpis 失败了！",err )
		return
	}

	return
}