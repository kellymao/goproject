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

func (m *Menu) TestMenu() []Menu {
	fmt.Println("test")

	var i []int
	fmt.Printf("%+v \n ",i)
	i = make([]int,0 )
	fmt.Printf("%+v \n ",i)

	var result []Menu

	result = append(result,Menu{Menuid:"999"})
	fmt.Println(result)
	return result
}

func (m *Menu) GetMenu (id string) (result []Menu,err error) {


	err = qmsql.DEFAULTDB.Where("authority_id=? and parent_id = 0 ",888).Find(&result).Error //delete_at 有值的不会显示出来
	if err!=nil {
		fmt.Println("ERROR:", err)
		return
	}

	//for _, rel := range result{
	//
	//	err = m.GetChildMenu(&rel)
	//	if err!=nil{
	//		return result,err
	//	}
	//
	//}

	for i:=0 ; i<len(result); i++ {

		err = m.GetChildMenu(&result[i])
		if err!=nil{
			fmt.Println("ERROR:", err)

			return
		}
	}

	//fmt.Printf("返回： %+v\n", result)
	return

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

	parent_menu.Children =  child_menu

	for i:=0 ; i<len(parent_menu.Children); i++ {

		err = m.GetChildMenu(&parent_menu.Children[i])
		if err!=nil{
			fmt.Println("ERROR:", err)

			return
		}
	}

	//fmt.Printf("返回childmenu ： %+v\n", child_menu)
	return nil 


}