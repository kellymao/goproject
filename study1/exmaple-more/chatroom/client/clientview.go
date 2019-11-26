package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"strings"
	"study1/exmaple-more/chatroom/common"
	"sync"
)

var wg sync.WaitGroup
var chatview Chatview

var log  *logrus.Logger


type  Chatview struct{

	client *ChatClient
	userid string

}


func (v Chatview) login() error{
	fmt.Println("请输入登录信息：")
	fmt.Println("请输入用户id:")
	userid := <- v.client.Msg_stdin
	fmt.Println("请输入密码：")
	pwd := <- v.client.Msg_stdin

	login_cmd := map[string]string{
		"userid":userid,
		"password":pwd,
	}

	login_data,err:=json.Marshal(login_cmd)
	if err!=nil{
		fmt.Println("login_data marshal err:",err)
	}

	msg:= common.Msginfo{
		Msgtype:"login",
		Msgtxt:string(login_data),
		Sendto:"server",
		From:userid,

	}

	data,err:=json.Marshal(msg)

	if err!=nil{
		fmt.Println("login marshal err:",err)
	}

	v.client.Msg_send <- string(data)

	if <-v.client.Login_status{
		fmt.Println("登录成功！",userid)
		v.userid = userid
	}else{
		fmt.Println("登录失败！",userid)
		return fmt.Errorf("登录失败,用户名或密码不对")
	}

	return nil

}

func (v Chatview) logout(){

	msg:= common.Msginfo{
		Msgtype:"logout",
		Msgtxt:"logout",
		Sendto:"server",
		From:v.userid,

	}
	data,err:=json.Marshal(msg)

	if err!=nil{
		fmt.Println("logout marshal err:",err)
	}

	v.client.Msg_send <- string(data)

}

func (v Chatview) logicview(ctx context.Context){

	fmt.Println("请输入聊天内容： 群发请输入 @all + 内容，单聊请输入 @userid + 内容, 退出输入logout")

	for {

		select {
			case content := <-v.client.Msg_stdin:
				if content=="logout"{
					v.logout()
					return
				}else if strings.HasPrefix(content,"@all"){

					msgtxt:=strings.TrimPrefix(content,"@all")
					msg:= common.Msginfo{
						Msgtype:"message",
						Msgtxt:msgtxt,
						Sendto:"all",
						From:v.userid,

					}
					data,err:=json.Marshal(msg)

					if err!=nil{
						log.Error("群聊 marshal err:",err)
					}

					log.Debug("传递message到Msg_send: ", string(data))
					v.client.Msg_send <- string(data)

				}else if strings.HasPrefix(content,"@"){

					msgtxt:=strings.Split(strings.TrimPrefix(content,"@")," ")

					msg:= common.Msginfo{
						Msgtype:"message",
						Msgtxt:msgtxt[1],
						Sendto:msgtxt[0],
						From:v.userid,

					}
					data,err:=json.Marshal(msg)

					if err!=nil{
						log.Error("单聊 marshal err:",err)
					}

					v.client.Msg_send <- string(data)
				}else{

					fmt.Println("格式输入错误！")
					fmt.Println("请输入聊天内容： 群发请输入 @all + 内容，单聊请输入 @userid + 内容, 退出输入logout")

				}

			case <-ctx.Done():
				return

		}

	}




}

func init_log(){

	log  = logrus.New()
	// 设置日志输出为os.Stdout
	log.Out = os.Stdout
	// 会记录info及以上级别 (warn, error, fatal, panic)
	log.SetLevel(logrus.DebugLevel) //InfoLevel

	log.WithFields(logrus.Fields{"socket": chatview.client.Chatsocket,"userid": chatview.userid})


	formatter := &logrus.TextFormatter{   // &logrus.JSONFormatter{}
		// 不需要彩色日志
		//DisableColors:   true,
		// 定义时间戳格式
		TimestampFormat: "2006-01-02 15:04:05",
		DisableTimestamp:false,

	}


	log.SetFormatter(formatter)
	//log.SetFormatter(&logrus.JSONFormatter{})
}

func init(){

	chatview =  Chatview{
		client:&ChatClient{
			Msg_stdin:make(chan string),
			Msg_send: make(chan string),
			Msg_recv: make(chan string),
			Login_status: make(chan bool),
		},
	}

	init_log()

}

func main(){

	fmt.Println("欢迎来到聊天室")

	if ok:=chatview.client.connect_server("127.0.0.1:10086");!ok{
		//fmt.Println("连接失败了")
		log.Error("connect_server连接失败了")
		return
	}
	log.Debug("socket 连接：",chatview.client.Chatsocket)

	ctx,cancal := context.WithCancel(context.Background())
	defer cancal()

	go chatview.client.handle_stdin()
	go chatview.client.handle_recv(ctx,cancal)
	go chatview.client.handle_message(chatview.userid,ctx)
	err:=chatview.login()
	if err!=nil{
		log.Error(err)
		return
	}



	chatview.logicview(ctx)


}


