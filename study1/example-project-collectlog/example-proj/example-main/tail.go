package main

import (

	"github.com/hpcloud/tail"
	"time"
)

import "github.com/astaxie/beego/logs"


func tailfile(collect *Collectconf) {

	//filename := ".\\my.log"
	tails, err := tail.TailFile(collect.Logfile, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		logs.Error("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {

		select {
		case msg, ok = <-tails.Lines :  // 管道被关闭 ok 为false
			if !ok {
				logs.Error("tail file close reopen, filename:%s\n", tails.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			//fmt.Println("msg:", msg)



			msgchan <- Message{

				topic:collect.Topic,
				msgtxt:msg.Text,
			}

		case 	<-collect.Exitchan:
			logs.Debug(collect.Logfile,"停止收集日志")
			return

		}
	}
}

