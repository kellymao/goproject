package main


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
	"database/sql"

)


var Db *sqlx.DB


func init(){

	db,err:=sqlx.Open("mysql","root:12345678@tcp(127.0.0.1:3306)/mydb1?charset=utf8mb4")

	if err!=nil{

		fmt.Println(err)
		panic(err)
	}
	Db = db

}

/*
	CREATE TABLE `userinfo` (
		`uid` INT(10) NOT NULL AUTO_INCREMENT,
		`username` VARCHAR(64)  DEFAULT NULL,
		`password` VARCHAR(32)  DEFAULT NULL,
		`department` VARCHAR(64)  DEFAULT NULL,
		`email` varchar(64) DEFAULT NULL,
		`login_time` datetime default null,
		`update_time` timestamp ,
		PRIMARY KEY (`uid`)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8

 */

func dml(){

	// insert
	result,err:=Db.Exec("insert into userinfo(uid,username,password,department,email,login_time) values (?,?,?,?,?,?)",1,"xiaomai","xxxxx","招聘","hr@129.com",time.Now().Format("2006-01-02 15:04:05"))

	if err!=nil{

		panic(err)
	}


	// 通过LastInsertId可以获取插入数据的id
	userId,err:= result.LastInsertId()
	// 通过RowsAffected可以获取受影响的行数
	rowCount,err:=result.RowsAffected()


	fmt.Println(userId)   //1
	fmt.Println(rowCount) //1
	fmt.Println(time.Now()) // 2019-10-10 11:16:23.370620694 +0800 CST m=+0.004333327

	// update

	result_u,err:=Db.Exec("update userinfo set username=? where uid =?","zizi",1)

	if err!=nil{
		fmt.Println("update err",err)
	}else{


		userId,_= result_u.LastInsertId()

		rowCount,_=result_u.RowsAffected()

		fmt.Println(userId)     // 0
		fmt.Println(rowCount)   // 1


	}


	result_del,err:=Db.Exec("delete from userinfo ")

	if err!=nil{

		fmt.Println("delete err",err)
	}else{
		rowCount,_=result_del.RowsAffected()
		fmt.Println(rowCount)  // 1

	}
}

func query(){

	result,err:=Db.Query("select uid,username,login_time,update_time from userinfo")

	if err!=nil{

		fmt.Println("query error")
		panic(err)
	}

	for result.Next() {


		/*
		type NullString struct {
		    String string
		    Valid  bool // Valid is true if String is not NULL
		}
		 */
		var uid int  //sql.NullInt64
		var username string
		var login_time sql.NullString
		var update_time sql.NullString

		err=result.Scan(&uid,&username,&login_time,&update_time)

		if err!=nil{

			panic(err)
		}

		fmt.Println(uid,username,login_time,update_time)


		/*

		1 zizi { false} {2019-10-10 11:33:18 true}
		2 xiaomai { false} {2019-10-10 11:32:59 true}


		mysql> select * from userinfo;
		+-----+----------+----------+------------+-------+------------+---------------------+
		| uid | username | password | department | email | login_time | update_time         |
		+-----+----------+----------+------------+-------+------------+---------------------+
		|   1 | zizi     | xxxx     |            |       | NULL       | 2019-10-10 11:33:18 |
		|   2 | xiaomai  | xxxx     |            |       | NULL       | 2019-10-10 11:32:59 |
		+-----+----------+----------+------------+-------+------------+---------------------+
		 */

	}


}


func queryx(){

	type person struct{

		Uid int                       `db:"uid"`
		Username string				  `db:"username"`
		Email string                  `db:"email"`
		Login_time sql.NullString     `db:"login_time"`
		Update_time sql.NullString    `db:"update_time"`

	}

	result,err:=Db.Queryx("select uid,username,email,login_time,update_time from userinfo")
	if err!=nil{

		panic(err)

	}

	for result.Next(){

		var p person
		err=result.StructScan(&p)
		if err!=nil{
			panic(err)
		}
		fmt.Println(p)
		fmt.Println(p.Username)
		/*
		{1 zizi  { false} {2019-10-10 11:33:18 true}}
		zizi
		{2 xiaomai  { false} {2019-10-10 11:32:59 true}}
		xiaomai
		 */

	}



}


func get_or_select(){

	type person struct{

		Uid int                       `db:"uid"`
		Username string				  `db:"username"`
		Email string                  `db:"email"`
		Login_time sql.NullString     `db:"login_time"`
		Update_time sql.NullString    `db:"update_time"`

	}

	var p []person


	err:=Db.Select(&p,"select uid,username,email,login_time,update_time from userinfo where uid=1")
	if err!=nil{

		panic(err)

	}

	fmt.Println(p) // [{1 zizi  { false} {2019-10-10 11:33:18 true}}]


	var s person
	err= Db.Get(&s,"select uid,username,email,login_time,update_time from userinfo where uid=1")

	if err!=nil{
		panic(err)
	}

	fmt.Println(s) // {1 zizi  { false} {2019-10-10 11:33:18 true}}


}


func prepare(){

	type person struct{

		Uid int                       `db:"uid"`
		Username string				  `db:"username"`
		Email string                  `db:"email"`
		Login_time sql.NullString     `db:"login_time"`
		Update_time sql.NullString    `db:"update_time"`

	}



	stmt,err:=Db.Preparex("select uid,username,email,login_time,update_time from userinfo where uid=? ")

	if err!=nil{
		panic(err)
	}

	var p []person
	err = stmt.Select(&p,1)

	fmt.Println(p)

	var s person
	err = stmt.Get(&p,2)
	fmt.Println(s)


}

func main(){


	//queryx()
	// get_or_select()

	prepare()

}



