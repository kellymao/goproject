package main


import "study1/example-chatroom/model"

var (

	mgr *model.Usermgr

)


func initUsermgr(){



		mgr:= model.NewUsermgr(pool)




}
