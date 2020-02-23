package main

import (
	"fmt"
	"study1/example-project/YunweiServer/config"
	"study1/example-project/YunweiServer/init/initRouter"
	"study1/example-project/YunweiServer/init/qmlog"
	"study1/example-project/YunweiServer/init/qmsql"
	"study1/example-project/YunweiServer/init/registTable"
	"net/http"
	"time"
)


func main() {
	qmlog.InitLog()
	registTable.RegistTable(qmsql.InitMysql(config.Dbconfig.Admin))
	defer qmsql.DEFAULTDB.Close()
	Router := initRouter.InitRouter()
	qmlog.QMLog.Info("服务器开启") // 日志测试代码

	s := &http.Server{
		Addr:           ":16666",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`前端文件运行地址:http://172.18.8.178:16668
						后端文件运行地址:http://172.18.8.178:16666
				`, s.Addr)
	_ = s.ListenAndServe()
}

func run(server *http.Server) {

}