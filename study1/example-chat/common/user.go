package common




type Friend struct{

	User_id []string


}

type User struct {

	Id string
	Pwd string
	Name string
	Status string
	Friends Friend




}

func (p *User) GetFriends(){


}
