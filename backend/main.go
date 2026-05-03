package main

import (
	"fmt"
	"log"

	"agriculturalAssistancePlatform/config"
	"agriculturalAssistancePlatform/database"
	"agriculturalAssistancePlatform/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	database.InitDB()

	if config.AppConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routers.SetupRouter()

	port := config.AppConfig.Server.Port
	addr := fmt.Sprintf(":%d", port)

	log.Printf("服务器启动在端口 %d", port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
