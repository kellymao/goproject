package main

import "net"
import "study1/example-chat/common"
import "encoding/json"
import "fmt"
import "errors"
import "strings"

type Handle struct{

	conn net.Conn
	userobj common.User


}

func (p *Handle) send_msg(msg common.Message){



	data,_:=json.Marshal(msg)
	fmt.Println("------",string(data))
	fmt.Printf("%v %p \n ",p.conn,p.conn)

	n,err:=p.conn.Write(data)
	fmt.Println("send byte",n)
	common.Check_err(err)
}

/*
func (p *Handle) send_resp(resp common.Resp){


	data,err:= json.Marshal(resp)
	if err!=nil{
		fmt.Println(err)
		return
	}
	p.send_msg(data)


}
*/

func (p *Handle) rev_msg() (msg common.Message, err error ) {


	var bufdata []byte

	var total,n int

	var buf [512]byte
	for {
		n, err = p.conn.Read(buf[0:])
		if err!=nil {

			fmt.Println("recive  data error",err)
			return
		}


		rAddr := p.conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))

		bufdata = append(bufdata, buf[0:n]...)
		total +=n

		break

	}



	json.Unmarshal(bufdata,&msg)

	return


}





func (p *Handle) login(msg common.Message) error {

	fmt.Println(msg.Msginfo)

	var user common.User = common.User{}

	err := json.Unmarshal([]byte(msg.Msginfo),&user)

	common.Check_err(err)

	var resp common.Resp =  common.Resp{}

	fmt.Println(*usermgr)

	user ,ok := usermgr.Checkuser(user.Id,user.Pwd)

	/*
	if user.Id  == "1" && user.Pwd == "1" {



		var resp common.Resp =  common.Resp{

			Status: true,
			Resptype: "login_resp",
			RespBody: msg.Msginfo,

		}

		data,err:= json.Marshal(resp)
		if err!=nil{


			return err
		}
		p.send_msg(data)

	}

	 */

	msg.Msgtype = "loginresp"

	if !ok{

		resp.Status = false

		data,_:= json.Marshal(resp)
		msg.Msginfo = string(data)
		p.send_msg(msg)

		return errors.New("user not exist.")


	}

	/*
	var notice common.Message  = common.Message{

		Msgtype:"notice",
		Msginfo: fmt.Sprintf("###### Notice: %s %s 已上线 ######",user.Id,user.Name) ,
		Msgsend: user.Id,
		Msgrecv:"all",
	}

	// 发送所有通知
	clientmgr.sendall(notice)
	*/

	// 在线列表加入本节点
	user.Status = "online"
	p.userobj = user

	fmt.Println("")
	fmt.Printf("加入 handle 到map： %v %p \n ",p,p)
	clientmgr.add(user.Id,p)



	data,_:=json.Marshal(user)

	resp.Status = true
	resp.RespBody = string(data)

	msgdata,_:= json.Marshal(resp)
	msg.Msginfo = string(msgdata)

	p.send_msg(msg)



	return nil


}

func (p *Handle) process()  {

	defer func(){
		clientmgr.del(p.userobj.Id)
	}()


for {

	msg ,err := p.rev_msg()
	if err!=nil{

		fmt.Println(err)
		return
	}

	msg_r1 := common.Message{

		Msgtype:"allmessage",
		Msginfo:p.userobj.Id + "说： " +"怎么回事啊" ,
		Msgsend: p.userobj.Id,
		Msgrecv: 	"all",
	}
	//data,_:=json.Marshal(msg_r)
	p.send_msg(msg_r1)

	switch msg.Msgtype {


	case common.Msgtype_login:

		continue
		err =p.login(msg)

	case common.Msgtype_getonlinelist:
		var keylist []string = make([]string,0 )

		for k,_ := range clientmgr.clients{

			keylist = append(keylist,k)

		}

		var msg_r common.Message = common.Message{

			Msgtype:"getonlinelist",
			Msginfo: strings.Join(keylist, " --- "),
			Msgsend: p.userobj.Id,
			Msgrecv: 	p.userobj.Id,
		}

		p.send_msg(msg_r)

	case common.Msgtype_getfriends:
		var msg_r common.Message = common.Message{

			Msgtype:"getfriends",
			Msginfo: strings.Join( p.userobj.Friends.User_id," --- "),
			Msgsend: p.userobj.Id,
			Msgrecv: 	p.userobj.Id,
		}

		p.send_msg(msg_r)

	case common.Msgtype_allmessage:


		fmt.Println("xxxxxxx")
		var msg_r common.Message = common.Message{

			Msgtype:"allmessage",
			Msginfo:p.userobj.Id + "说： " + msg.Msginfo ,
			Msgsend: p.userobj.Id,
			Msgrecv: 	"all",
		}


		 msg_r = common.Message{

			Msgtype:"allmessage",
			Msginfo:p.userobj.Id + "说： " +"怎么回事啊" ,
			Msgsend: p.userobj.Id,
			Msgrecv: 	"all",
		}
		//data,_:=json.Marshal(msg_r)
		p.send_msg(msg_r)

		for i := range clientmgr.clients{

			fmt.Println("这里面有没有东西")
			fmt.Printf("[%s]",i)
			clientmgr.clients[i].send_msg(msg_r)

			fmt.Println(p.conn,clientmgr.clients[i].conn )
			fmt.Printf("%p %p\n",p.conn,clientmgr.clients[i].conn)
			p.send_msg(msg_r)
			p.send_msg(msg_r)
			clientmgr.clients[i].send_msg(msg_r)

		}

		//clientmgr.sendall(msg_r)

	case common.Msgtype_singlemessage:

		var msg_r common.Message = common.Message{

			Msgtype:"singlemessage",
			Msginfo:p.userobj.Id + "说： " + msg.Msginfo ,
			Msgsend: p.userobj.Id,
			Msgrecv: msg.Msgrecv,
		}

		to_handle :=clientmgr.clients[msg.Msgrecv]
		to_handle.send_msg(msg_r)

	default:
		return

	}

	if err!=nil{


		fmt.Println(err)

	}

}




}
