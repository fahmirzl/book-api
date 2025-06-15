package routers

import (
	"book-api/handlers"
	"book-api/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(rg *gin.RouterGroup) {
	category := rg.Group("/categories")
	category.GET("/", middlewares.Auth(), handlers.CategoryIndex)
	category.POST("/", middlewares.Auth(), handlers.CategoryStore)
	category.GET("/:id", middlewares.Auth(), handlers.CategoryFind)
	category.PUT("/:id", middlewares.Auth(), handlers.CategoryUpdate)
	category.DELETE("/:id", middlewares.Auth(), handlers.CategoryDestroy)
	category.GET("/:id/books", middlewares.Auth(), handlers.CategoryWithBook)
}