package config

import (
	"log"
	"strings"
)

var AppConfBO = new(appConf)

type appConf struct {
	Port, CpName, Env, LogPath, LogFile, LogLevel, MysqlHost, MysqlPort, MysqlUser, MysqlPass, MysqlDatabase, MysqlMaxConn, MysqlMaxIdle, QueryTimeout string
}

func (c *appConf) InitConfig(filePath string) {
	log.Printf("[file change call back, init start]:%s", filePath)
	InitConfig(filePath)
}

func InitConfig(filePath string) {
	err := LoadUmshal(filePath, Yaml, AppConfBO, AppConfBO)
	if err != nil {
		log.Fatal(err)
	}

	if strings.TrimSpace(AppConfBO.Port) == "" {
		log.Fatal("【启动失败】app.yaml配置文件port为空")
	}

	if strings.TrimSpace(AppConfBO.CpName) == "" {
		log.Fatal("【启动失败】app.yaml配置文件cpName为空")
	}

	if strings.TrimSpace(AppConfBO.Env) == "" {
		log.Fatal("【启动失败】app.yaml配置文件env为空")
	}

	if strings.TrimSpace(AppConfBO.LogPath) == "" {
		log.Fatal("【启动失败】app.yaml配置文件logPath为空")
	}

	if strings.TrimSpace(AppConfBO.LogFile) == "" {
		log.Fatal("【启动失败】app.yaml配置文件logFile为空")
	}

	if strings.TrimSpace(AppConfBO.LogLevel) == "" {
		log.Fatal("【启动失败】app.yaml配置文件logFile为空")
	}

	if strings.TrimSpace(AppConfBO.MysqlHost) == "" {
		log.Fatal("【启动失败】app.yaml配置文件mysqlHost为空")
	}

	if strings.TrimSpace(AppConfBO.MysqlPort) == "" {
		log.Fatal("【启动失败】app.yaml配置文件mysqlPort为空")
	}

	if strings.TrimSpace(AppConfBO.MysqlUser) == "" {
		log.Fatal("【启动失败】app.yaml配置文件mysqlUser为空")
	}

	if strings.TrimSpace(AppConfBO.MysqlPass) == "" {
		log.Fatal("【启动失败】app.yaml配置文件mysqlPass为空")
	}

	if strings.TrimSpace(AppConfBO.QueryTimeout) == "" {
		log.Fatal("【启动失败】app.yaml配置文件queryTimeout地址为空")
	}

}
