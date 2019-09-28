package main
import "github.com/garyburd/redigo/redis"
import "fmt"
import "study1/example-chat/common"
import "encoding/json"
import "errors"


var usermgr *Usermgr = &Usermgr{}


func init_usermgr()  {

	usermgr  = &Usermgr{

		pool: pool,

	}

	/*
	usermgr.AddUser("11","{\"Id\":\"11\",\"Pwd\":\"11\",\"Name\":\"张三\",\"Status\":\"\",\"Friends\":{\"User_id\":[\"12\",\"13\",\"14\"]}}")
	usermgr.AddUser("12","{\"Id\":\"12\",\"Pwd\":\"12\",\"Name\":\"赵高\",\"Status\":\"\",\"Friends\":{\"User_id\":[\"13\"]}}")

	usermgr.AddUser("13","{\"Id\":\"13\",\"Pwd\":\"13\",\"Name\":\"美丽\",\"Status\":\"\",\"Friends\":{\"User_id\":[\"11\",\"14\"]}}")

	usermgr.AddUser("14","{\"Id\":\"14\",\"Pwd\":\"14\",\"Name\":\"真难\",\"Status\":\"\",\"Friends\":{\"User_id\":[\"12\"]}}")
	*/
	return

}


type Usermgr struct{

	pool *redis.Pool


}

func (p *Usermgr) Checkuser(id string,pwd string)(user common.User, flag bool){

	// 检查用户是否存在
	user,err:= p.Getuser(id)
	if err!=nil{
		flag = false
		return
	}

	//检查密码是否正确

	if user.Pwd == pwd {

		flag = true
	}


	return


}

func (p *Usermgr) Getuser(id string) (user common.User,err error ){

	conn := p.pool.Get()
	defer conn.Close()

	result,err := redis.String(conn.Do("HGet","user",id))
	if err!=nil{

		if err == redis.ErrNil {
			err = errors.New("user not exist")
			fmt.Println(err)
		}
		return
	}

	err = json.Unmarshal([]byte(result),&user)

	if err!=nil{

		fmt.Println(err)
		return
	}
	return

}

func (p *Usermgr) AddUser(id string,data string){


	conn := p.pool.Get()
	defer conn.Close()
	conn.Do("HSet","user",id,data)


}

func (p *Usermgr) DelUser(id string){

	conn := p.pool.Get()
	defer conn.Close()
	conn.Do("HDel","user",id)

}

func (p *Usermgr) GetAllonline() string{

	conn := p.pool.Get()
	defer conn.Close()
	conn.Do("HGet","user")


	return ""

}


