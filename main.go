package main

import (
	"drawerBackend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.Routers(r)

	_ = r.Run(":8080")
}
