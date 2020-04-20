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

	rows,err := qmsql.DEFAULTDB.Model(&Role_to_menu{}).Select("menuid").Where("roleid=? ",id).Rows() // 先取出该角色可以访问的menu_id


	if err!=nil{
		return
	}
	defer rows.Close()


	var scan_id string

	var ids []string


	for rows.Next() {
		_ = rows.Scan(&scan_id)
		ids = append(ids,scan_id)
	}

	err = qmsql.DEFAULTDB.Where("menuid in (?) and parent_id = 0 ",ids).Find(&result).Error //delete_at 有值的不会显示出来
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

		err = m.GetChildMenu(&result[i],ids)
		if err!=nil{
			fmt.Println("ERROR:", err)

			return
		}
	}

	return

}



func (m *Menu) GetChildMenu(parent_menu *Menu,ids []string)( err error)  {

	var child_menu []Menu
	//err = qmsql.DEFAULTDB.Where("authority_id=? and parent_id = ? ",parent_menu.AuthorityId , parent_menu.Menuid).Find(&child_menu).Error
	err = qmsql.DEFAULTDB.Where("menuid in (?) and parent_id = ? ", ids,parent_menu.Menuid).Find(&child_menu).Error

	if err!=nil{
		return
	}

	//fmt.Println("========",parent_menu.Menuid)
	//fmt.Printf("%+v\n",ids)

	if len(child_menu) == 0 {

		return nil
	}

	parent_menu.Children =  child_menu

	for i:=0 ; i<len(parent_menu.Children); i++ {

		err = m.GetChildMenu(&parent_menu.Children[i],ids)
		if err!=nil{
			fmt.Println("ERROR:", err)

			return
		}
	}

	//fmt.Printf("返回childmenu ： %+v\n", child_menu)
	return nil 


}