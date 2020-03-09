package dbModel

import (
	"fmt"
	"study1/example-project/YunweiServer/init/qmsql"
)

type Role_to_menu struct {

	Roleid string
	Menuid  string `json:"menuId"`
	Parentid string `json:"parentid"`



}


type Role_menu_tree struct{
	Role_to_menu
	Menutitle string `json:"title"`
	Child  []Role_menu_tree  `json:"children"`

}



func (m *Role_to_menu) Save (roleid string,menu []Role_to_menu) (err error) {


	for i:=0; i < len(menu); i++ {

		err = qmsql.DEFAULTDB.Save(&menu[i]).Error
		if err!=nil {
			fmt.Println("ERROR:", err)
			return
		}
	}



	return

}

func (m *Role_to_menu)  Get_restree(roleid string) (restree []Role_menu_tree , err error){


	var result []Role_to_menu

	err = qmsql.DEFAULTDB.Where("Roleid=? and parent_id = 0 ",roleid).Find(&result).Error //delete_at 有值的不会显示出来
	if err!=nil {
		fmt.Println("ERROR:", err)
		return
	}



	for i:=0;i<len(result);i++{

		res:= Role_menu_tree{
			result[i],

		}

	}



}

func get_title(Menuid string) (title string ,err error ){
	err = qmsql.DEFAULTDB.Model(&Menu{}).Select("name").Where("Menuid=?  ",Menuid).Find(&title).Error //delete_at 有值的不会显示出来

	if err!=nil{
		return
	}

	return

}


func get_child(roleid string,Menuid string) (restree Role_menu_tree , err error) {
	err = qmsql.DEFAULTDB.Where("Roleid=? and parent_id = 0 ",roleid).Find(&result).Error //delete_at 有值的不会显示出来



}