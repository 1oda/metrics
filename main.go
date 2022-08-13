package main

import (
	"fmt"
	"metrics/router"
	"metrics/utils/config"
	"metrics/utils/db"
	logs "metrics/utils/log"

	"github.com/gin-gonic/gin"
)

const (
	configFile = "./config/app.yml"
)

func init() {

	// 初始化全局配置
	config.InitConfig(configFile)

	// 初始化全局日志配置
	logs.InitLog()

	// 初始化全局数据库配置
	db.InitDB()
}

func main() {

	r := gin.Default()

	router.Load(r)

	logs.Info(fmt.Sprintf("server running on 0.0.0.0:%s", config.AppConfBO.Port))
	_ = r.Run(fmt.Sprintf("0.0.0.0:%s", config.AppConfBO.Port))
}
