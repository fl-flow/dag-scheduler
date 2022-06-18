package httpserver

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/fl-flow/dag-scheduler/docs"
	"github.com/fl-flow/dag-scheduler/http_server/app"
	"github.com/fl-flow/dag-scheduler/http_server/http/middleware"
)

func Run(ip string, port int) {
	ginApp := gin.Default()
	ginApp.Use(middleware.AuthMiddleware)
	ginApp.GET("/version", func(c *gin.Context) {
    c.String(http.StatusOK, "1.0.0")
  })
	ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginApp.POST("/v1/job/", app.JobCreate)
	ginApp.GET("/v1/job/", app.JobList)

	ginApp.Run(fmt.Sprintf("%v:%d", ip, port))
}
