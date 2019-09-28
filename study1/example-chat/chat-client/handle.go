package main



import "fmt"
import "study1/example-chat/common"
import "encoding/json"
import "net"
import "io"
import "time"
import "errors"

type Handle struct{



	conn net.Conn
	talk_mode  int
}



func (p *Handle) run() {

	p.login()


	// 消息通道有消息写入，就发送到服务器
	go func(){

		for {
			msg := <-chan_write

			data, err := json.Marshal(msg)
			common.Check_err(err)
			p.send_msg(data)
		}


	}()

	go func() {
		for {

			msg, err := p.rev_msg()
			if err != nil {
				fmt.Println(err)
				continue
			}


			fmt.Println("客户端收到了")
			fmt.Println(msg)

			switch msg.Msgtype {

			case common.Msgtype_loginresp:

				var resp common.Resp = common.Resp{}

				err = json.Unmarshal([]byte(msg.Msginfo), &resp)
				common.Check_err(err)

				if ! resp.Status {

					fmt.Println("登录失败！请重新登录")
					p.login()

				} else {

					fmt.Println("登录验证通过")

					err = json.Unmarshal([]byte(resp.RespBody), &userinfo)
					common.Check_err(err)

					p.talk()

				}


			case common.Msgtype_notice:
				fmt.Println("通知：",msg.Msginfo)

			case common.Msgtype_allmessage:

				if p.talk_mode ==1 {

					fmt.Println(msg.Msginfo)
				}

			case common.Msgtype_singlemessage:
				if p.talk_mode ==2 {

					fmt.Println(msg.Msginfo)
				}

			case common.Msgtype_getonlinelist:



				fmt.Println("在线用户列表：",msg.Msginfo)
			case common.Msgtype_getfriends:
				fmt.Println("你的好友列表：",msg.Msginfo)


			default:
				continue

			}

		}
	}()

	<-chan_exit


}

func (p *Handle) send_msg(data []byte){


	_,err:=p.conn.Write(data)
	common.Check_err(err)
}

func (p *Handle) rev_msg() (common.Message, error ) {


	var msg common.Message = common.Message{}
	var err error
	var bufdata []byte

	var total,n int

	var buf [512]byte
	for {
		n, err = p.conn.Read(buf[0:])
		if err!=nil {

			if err == io.EOF {

				err = errors.New("recive  data EOF")
				break


			}

			fmt.Println("recive  data error",err)
			err = errors.New("recive  data error")

			break
		}


		rAddr := p.conn.RemoteAddr()
		fmt.Println("Receive from server: ", rAddr.String(), string(buf[0:n]))

		bufdata = append(bufdata, buf[0:n]...)
		total +=n

		/*
		if err == io.EOF {

			err = errors.New("recive  data EOF")
			break


		}

		 */

		if n < 512 {              // 小于512 说明接受的内容已经接受完毕，可以退出循环

			break
		}




	}



	json.Unmarshal(bufdata,&msg)

	return msg,err


}

func(p *Handle) login(){

	// 从终端输入用户名密码
	prompt()

	var user common.User = common.User{
		Id: userid,
		Pwd:pwd,


	}

	user_str,err := json.Marshal(user)
	common.Check_err(err)



	var msg common.Message  = common.Message{

		Msgtype : "login",
		Msginfo: string(user_str),


	}


	data,err := json.Marshal(msg)

	common.Check_err(err)


	// 发送数据到服务器

	p.send_msg(data)

	return




}

func(p *Handle) talk(){


	p.process_talk()


}


func (p *Handle) process_talk(){


	fmt.Println("选择哪个聊天室，1：公共聊天室 2. 找好友单聊")

	var sel string
	fmt.Scanln(&sel)

	if sel == "1" {
		p.talk_mode = 1

		fmt.Printf("%s : 欢迎 %s 来到公用聊天室\n",time.Now().Format("02/1/2006 15:04"),userinfo.Name)

		// 获取在线人的列表

		var msg common.Message =  common.Message{

			Msgtype:"getonlinelist",
			Msginfo:"",
		}


		chan_write<-msg


		 go process_stdin("allmessage")

	}

	if sel == "2" {

		p.talk_mode = 2

		fmt.Printf("%s : 欢迎 %s 来单聊\n",time.Now().Format("02/1/2006 15:04"),userinfo.Name)

		// 显示好友的列表

		data,err:= json.Marshal(userinfo.Friends.User_id)
		common.Check_err(err)

		var msg common.Message =  common.Message{}

		msg.Msgtype = "getfriends"
		msg.Msginfo = string(data)



		chan_write<-msg


		go process_stdin("singlemessage")


	}

}


//func (p *Handle)