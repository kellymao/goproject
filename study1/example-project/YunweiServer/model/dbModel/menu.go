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

func (m *Menu) TestMenu()[]Menu{
	var rel []Menu 
	return rel
}

func (m *Menu) GetMenu (id string) ([]Menu,error) {

	var result []Menu

	err := qmsql.DEFAULTDB.Where("authority_id=? and parent_id = 0 ",id).Find(&result).Error //delete_at 有值的不会显示出来
	if err!=nil {
		return result,err
	}

	//for _, rel := range result{
	//
	//	err = m.GetChildMenu(&rel)
	//	if err!=nil{
	//		return result,err
	//	}
	//
	//}


	//fmt.Printf("返回： %+v\n", result)
	return result,nil

}



func (m *Menu) GetChildMenu(parent_menu *Menu)( err error)  {

	var child_menu []Menu
	err = qmsql.DEFAULTDB.Where("authority_id=? and parent_id = ? ",parent_menu.AuthorityId , parent_menu.Menuid).Find(&child_menu).Error

	if err!=nil{
		return
	}

	if len(child_menu) == 0 {

		return nil
	}
	for _, rel:= range child_menu{
		parent_menu.Children = append(parent_menu.Children,rel)
		err = m.GetChildMenu(&rel)
		if err!=nil {
			return
		}

	}
	fmt.Printf("返回childmenu ： %+v\n", child_menu)
	return nil 


}