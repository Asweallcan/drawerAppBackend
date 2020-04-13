package routers

import (
	"drawerBackend/controllers"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.GET("/", controllers.Home)
}
