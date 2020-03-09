package dbModel
import "github.com/jinzhu/gorm"

import (
	"fmt"
	"study1/example-project/YunweiServer/init/qmsql"
)

type Role_to_menu struct {

	gorm.Model
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

	// 得到了roleid = 888 的所有菜单权限。要求是要展示所有的菜单。有权限的打打钩

	var result []Role_to_menu

	err = qmsql.DEFAULTDB.Select("roleid,menuid,parentid").Where("roleid=? and parentid = 0 ",roleid).Find(&result).Error //delete_at 有值的不会显示出来
	if err!=nil {
		fmt.Println("ERROR:", err)
		return
	}



	for i:=0;i<len(result);i++{

		title,_:= get_title(result[i].Menuid)
		child,_:= get_child(&result[i])
		res:= Role_menu_tree{
			result[i],
			title,
			child,


		}

		restree =append(restree,res)


	}

	return



}

func get_title(Menuid string) (title string ,err error ){

	type T struct{
		Name string
	}

	var rel  = T{}
	fmt.Println(Menuid)
	err = qmsql.DEFAULTDB.Model(&Menu{}).Select("name").Where("menuid=?  ",Menuid).Scan(&rel).Error //delete_at 有值的不会显示出来

	if err!=nil{
		fmt.Println("获取title失败！",err)
		return
	}

	fmt.Printf("%+v \n",rel)
	return rel.Name,nil

}


func get_child(r *Role_to_menu) (restree []Role_menu_tree , err error) {

	var result []Role_to_menu

	err = qmsql.DEFAULTDB.Select("roleid,menuid,parentid").Where("roleid=? and parentid = ? ",r.Roleid,r.Menuid).Find(&result).Error //delete_at 有值的不会显示出来

	if err!=nil{
		fmt.Println("获取child失败！",err)
		return

	}

	if len(result) == 0 {
		return
	}

	for i:=0;i<len(result);i++{

		title,_:= get_title(result[i].Menuid)
		child,_:= get_child(&result[i])
		res:= Role_menu_tree{
			result[i],
			title,
			child,

		}

		restree =append(restree,res)

	}

	return



}



type SqlRes struct {
	Checked        string  `json:"checked"`
	Menuid string          `json:"menuid"`
	Parentid       string  `json:"parentid"`
	Name          string  `json:"title"`
	Child  []SqlRes  `json:"children"`
}

func (m *Role_to_menu) Get_role_menutree(roleid string)([]SqlRes,error){

	/*
	select (case  roleid is NULL when true  then 0 else 1 end ) as checked ,menus.id,menus.parent_id,name  from menus left join role_to_mee_to_menus  on menus.id = role_to_menus.menuid and  role_to_menus.roleid = '888'  ;
	+---------+----+-----------+--------------+
	| checked | id | parent_id | name         |
	+---------+----+-----------+--------------+
	|       1 |  2 | 0         | 系统管理     |
	|       1 |  3 | 0         | 导航二       |
	|       1 |  4 | 0         | 导航三       |
	|       1 |  5 | 0         | 导航四       |
	|       1 |  1 | 0         | 404          |
	|       1 |  6 | 2         | 主页         |
	|       1 |  7 | 2         | 表格         |
	|       1 |  8 | 2         | 表单         |
	|       1 |  9 | 2         | 用户管理     |
	|       1 | 10 | 2         | 角色管理     |
	|       0 | 11 | 2         | 编辑         |
	+---------+----+-----------+--------------+
	 */
	var result []SqlRes
	var sqlRes SqlRes
	rows, err  := qmsql.DEFAULTDB.Raw("select (case  roleid is NULL when true  then 0 else 1 end ) as checked ,menus.id,menus.parent_id,name  from menus left join role_to_menus  on menus.id = role_to_menus.menuid and  role_to_menus.roleid =   ? where menus.parent_id =0", roleid).Rows()
	//err:= row.Find(&sqlRes).Error

	if err!=nil{
		fmt.Println("获取menutree 失败！	")
		return nil,err
	}

	defer rows.Close()

	var checked , menuid, parentid , name string
	for rows.Next() {
		//rows.Scan(&sqlRes)
		rows.Scan(&checked,&menuid,&parentid,&name)

		sqlRes = SqlRes{
			checked,menuid,parentid,name,nil,
		}

		err = Get_child_menutree( roleid,&sqlRes)
		result = append(result,sqlRes)
	}

	fmt.Printf("%+v \n",result)
	return result,nil
	//row := qmsql.DEFAULTDB.Raw("SELECT apis.path,api_authorities.authority_id,api_authorities.api_id,apis.id FROM apis INNER JOIN api_authorities ON api_authorities.api_id = apis.id 	WHERE apis.path = ? AND	api_authorities.authority_id = ?", c.Request.RequestURI, claims.AuthorityId)
	//err = row.Scan(&sqlRes).Error

}


func Get_child_menutree(roleid string,r *SqlRes) error {

	var child_sqlRes SqlRes
	rows, err  := qmsql.DEFAULTDB.Raw("select (case  roleid is NULL when true  then 0 else 1 end ) as checked ,menus.id,menus.parent_id,name  from menus left join role_to_menus  on menus.id = role_to_menus.menuid and  role_to_menus.roleid =   ? where menus.parent_id =? ", roleid, r.Menuid).Rows()
	if err!=nil{
		fmt.Println("获取child menutree 失败！	",err)
		return err
	}

	defer rows.Close()

	var checked , menuid, parentid , name string

	for rows.Next() {
		rows.Scan(&checked,&menuid,&parentid,&name)

		child_sqlRes = SqlRes{
			checked,menuid,parentid,name,nil,
		}
		err = Get_child_menutree( roleid,&child_sqlRes)
		r.Child = append(r.Child,child_sqlRes)
	}

	return nil
}