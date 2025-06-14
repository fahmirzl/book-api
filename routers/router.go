package routers

import (
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	CategoryRouter(api)
	BookRouter(api)

	return r
}
