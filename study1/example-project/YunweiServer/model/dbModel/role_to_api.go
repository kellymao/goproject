package dbModel

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"study1/example-project/YunweiServer/init/qmsql"
)

type Role_to_api struct {

	gorm.Model
	Roleid string `json:"roleid"`
	Apiid  string `json:"apiid"`
}







func (m *Role_to_api) GetApi_list(roleid string)(ids []string,err error){


	//var apiList []Api
	rows, err  := qmsql.DEFAULTDB.Model(&Role_to_api{}).Select("apiid").Where("`roleid` = ?" , roleid).Rows()

	if err!=nil{
		return
	}

	defer rows.Close()

	var id string
	for rows.Next() {
		rows.Scan(&id)
		ids = append(ids,id)
	}

	return

}


func (m *Role_to_api) SaveApi_list(roleid string, ids []int )(err error) {

	tx := qmsql.DEFAULTDB.Begin()
	tx.Delete(&Role_to_api{},"roleid = ? ",roleid)


	for _,id := range ids{

		role_to_api := &Role_to_api{}
		role_to_api.Roleid = roleid
		role_to_api.Apiid = strconv.Itoa(id)


		err = tx.Save(role_to_api).Error
		if err!=nil{
			tx.Rollback()
			return
		}

	}


	tx.Commit()
	return
}

