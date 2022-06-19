package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/fl-flow/dag-scheduler/etc"
	_ "github.com/fl-flow/dag-scheduler/docs"
	"github.com/fl-flow/dag-scheduler/http_server/v1"
	"github.com/fl-flow/dag-scheduler/http_server/http/middleware"
)

func Run() {
	ginApp := gin.Default()

	ginApp.Use(middleware.AuthMiddleware)

	ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1.RegisterRouter(ginApp.Group("api/v1"))
	
	ginApp.Run(fmt.Sprintf("%v:%d", etc.SchedulerIp, etc.SchedulerPort))
}
