package app

import (
	"mychatgpt/app/controller"
	_ "mychatgpt/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Server     *gin.Engine
	Controller *controller.Controller
}

func (p *App) Start() {
	// 服务容灾recover
	p.Server.Use(p.Controller.ServeRecover)

	rootGroup := p.Server.Group("/")
	{
		rootGroup.GET("/", p.Controller.Index)
		rootGroup.GET("/index", p.Controller.Index)
		rootGroup.GET("/test/", p.Controller.Test)
		rootGroup.POST("/chatgpt", p.Controller.ChatGPT)

		//执行命令 swag init 初始化swagger
		//http://localhost:9090/swagger/index.html
		rootGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	logrus.Println("My ChatGPT 启动成功...")
	p.Server.Run(":9090")
}
