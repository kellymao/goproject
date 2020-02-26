package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"study1/example-project/YunweiServer/model/dbModel"
	"log"
)

var DEFAULTDB *gorm.DB

type Admin struct {
	Username string
	Password string
	Path     string
	Dbname   string
	Config   string
}

//初始化数据库并产生数据库全局变量
func InitMysql(admin Admin) *gorm.DB {
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		log.Printf("DEFAULTDB数据库启动异常%S", err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(10)
		DEFAULTDB.DB().SetMaxIdleConns(100)
	}
	return DEFAULTDB
}


// deleted_at = null ， 有值的goorm以为此记录被删除，select出不来结果。搞半天
func main(){

	cfg := Admin{
		Username : "root",
		Password: "12345678",
		Path: "127.0.0.1:3306",
		Dbname: "yyserver",
		Config: "charset=utf8&parseTime=True&loc=Local",
	}
	InitMysql(cfg)
	//
	//var result []dbModel.Menu
	//////err := DEFAULTDB.Find(&result).Error
	//////fmt.Println(err)
	//err := DEFAULTDB.Where("authority_id=? and parent_id = 0 ",888).Find(&result).Error
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%+v\n",result)
	//
	//
	//var users []dbModel.User
	//
	//_ = DEFAULTDB.Find(&users).Error
	//fmt.Println(users)

	aaa := &dbModel.Menu{AuthorityId: "888", Menuid: "2"}
	_, _ = aaa.GetMenu(888)
	//_ = aaa.GetChildMenu(&dbModel.Menu{AuthorityId: "888", Menuid: "2"})
	fmt.Printf("%T \n ",aaa)
	//aaa.TestMenu()

}