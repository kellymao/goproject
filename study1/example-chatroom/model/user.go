package model

const (

	UserStatusOnline = 1
	UserStatusOffline = itoa

)


type User struct{


	Userid int `json:"user_id"`
	Passwd string `json:"passwd"`
	Nick string `json:"nick"`
	Sex string `json:"json"`
	Header string `json:"header"`
	Lastlogin string `json:"lastlogin"`
	Status string `json:"status"`


}
