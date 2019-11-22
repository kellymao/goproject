package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appconf.log_path
	config["level"] = convertLogLevel(appconf.log_level)

	configStr, err := json.Marshal(config)
	//fmt.Println(string(configStr))
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))


	return
}
