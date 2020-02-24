package dbModel

import (
	"fmt"
	"study1/example-project/YunweiServer/init/qmsql"
)




type Menu struct {

	BaseMenu
	Menuid  string `json:"menuId"`
	AuthorityId string `json:"-"`
	Children    []Menu `json:"children"`


}


func (m *Menu) GetMenu (id string) ([]*Menu,error) {

	var result []*Menu = make([]*Menu ,0 )

	err := qmsql.DEFAULTDB.Where("authority_id=? and parent_id = 0 ",id).Find(&result).Error
	if err!=nil {
		return result,err
	}

	for _, rel := range result{

		err = rel.GetChildMenu()
		if err!=nil{
			return result,err
		}

	}


	fmt.Printf("%+v\n", result)
	return result,nil

}



func (m *Menu) GetChildMenu() error {

	var childmenu []Menu
	err := qmsql.DEFAULTDB.Where("authority_id=? and parent_id = ? ",m.AuthorityId , m.Menuid).Find(&childmenu).Error

	if err!=nil{
		return err 
	}

	m.Children = childmenu
	return nil 


}