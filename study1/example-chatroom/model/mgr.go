package model

import "github.com/garyburd/redigo/redis"
import "time"

var (

	UserTable = "users"

)
type Usermgr struct{

	pool *redis.Pool



}


func Newusermgr(pool *redis.Pool) *Usermgr{


	mgr:=&Usermgr{

		pool:pool,
	}

	return mgr


}

func (p *Usermgr) getUser(conn redis.Conn,id string) (*User,error){


	result,err:=redis.String(conn.Do("HGet",UserTable,fmt.Sprintf("%d",id)))

	if err!=nil{

		return nil,err
	}

	user :=&User{}
	err = json.Unmarshal([]byte(result),user)

	if err!=nil{
		return nil,err

	}

	return user ,nil




}

func (p *Usermgr)Login(id int,passwd string)(*User,error){

	conn:=p.pool.Get()
	defer conn.Close()


	user,err:=p.getUser(conn,id)
	if err!=nil{

		return nil,err
	}

	if user.Userid !=id || user.Passwd!=passwd {

		err = ErrUserNotExist
	}

	user.Status = UserStatusOnline
	user.Lastlogin = fmt.Sprintf("%v",time.Now())

	return user,nil
}


func (p *Usermgr) Register(user *User) error {


	conn:=p.pool.Get()
	defer conn.Close()

	if user == nil{

		fmt.Println("invalid user")
		err := ErrInvalidParams
		return err
	}

	user,err:= p.getUser(conn,user.Userid)

	if err==nil{

		err = ErrUserExist

		return err
	}

	data,err:= json.Marshal(user)

	if err!=nil{

		return err
	}

	_,err=conn.Do("HSet",UserTable,fmt.Sprintf("%id",user.Userid),string(data))

	if err!=nil{

		return err
	}

	return nil

}