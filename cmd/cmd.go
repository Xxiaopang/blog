package cmd

import (
	"go.uber.org/zap"

	"project/config"
	"project/database"
	"project/logs"
	"project/router"
)

func init() {
	config.InitConfig()
	database.InitDB()
	logs.InitLogger()
}

func registerRouter() {
	r := router.Init()

	r.Run(":" + config.App.Port)
	zap.S().Debugf("启动服务器，端口为：%d", config.App.Port)

}

func Execute() {
	registerRouter()
}
