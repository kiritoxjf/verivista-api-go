package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"verivista-api-go/config"
	"verivista-api-go/database"
	"verivista-api-go/handlers"
	"verivista-api-go/interfaces"
	"verivista-api-go/logger"
)

// VeriVistaConfig 配置数据
var VeriVistaConfig interfaces.Config

func main() {
	err := logger.SetLogOutPut()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	VeriVistaConfig, err = config.GetConfig()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	err = database.InitDB(VeriVistaConfig.DB)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	// 创建一个Gin路由器
	r := gin.Default()

	r.GET("/api/poem", handlers.PoemHandler)
	r.GET("/api/icp", handlers.ICPHandler)

	err = r.Run("127.0.0.1:8081")
	if err != nil {
		logrus.Errorln("[Error gin running]: ", err)
	}
}
