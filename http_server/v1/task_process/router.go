package taskprocess

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
	Router.POST("/", TaskProcess)
}
