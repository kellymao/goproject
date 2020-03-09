package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"log"
	"study1/example-project/YunweiServer/model/dbModel"

	"github.com/dgrijalva/jwt-go"



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

/*
// deleted_at = null ， 有值的goorm以为此记录被删除，select出不来结果。搞半天
func test_db(){

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
	_, _ = aaa.GetMenu("888")
	//_ = aaa.GetChildMenu(&dbModel.Menu{AuthorityId: "888", Menuid: "2"})
	fmt.Printf("%T \n ",aaa)
	//aaa.TestMenu()

}


 */



type JWT struct {
	SigningKey []byte
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}


var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "qmPlus"
)

type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	NickName    string
	AuthorityId string
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

//获取token
func GetSignKey() string {
	return SignKey
}
func test_token(){
	j := NewJWT()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEsIk5pY2tOYW1lIjoi6LaF57qn566h55CG5ZGYIiwiQXV0aG9yaXR5SWQiOiI4ODgiLCJleHAiOjE1ODMzOTMyNjksImlzcyI6InFtUGx1cyIsIm5iZiI6MTU4Mjc4NzQ2OX0.u1nEV_uz6HEavC2Yn5bWv5Qgz73rpfKpSUg5_toNJ90"
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == TokenExpired {
			fmt.Println(err)
		}

		fmt.Println("other:",err)
		return
	}

	fmt.Println(claims)
}

func test_restree(){
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

	aaa := &dbModel.Role_to_menu{}
	result, _ := aaa.Get_restree("888")
	fmt.Printf("%+v \n ",result)

}

func main(){

	//test_token()
	test_restree()


}