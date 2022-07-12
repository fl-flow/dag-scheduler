package task

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
	Router.POST("/torun/", TaskRunning)

	Router.POST("/cancel/", TaskCancelView)
}
