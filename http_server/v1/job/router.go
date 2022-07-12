package job

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
	Router.POST("", JobCreateView)
	Router.GET("", JobListView)
}
