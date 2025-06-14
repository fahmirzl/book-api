package routers

import (
	"book-api/handlers"
	"book-api/middlewares"

	"github.com/gin-gonic/gin"
)

func BookRouter(rg *gin.RouterGroup) {
	book := rg.Group("/books")
	book.GET("/", middlewares.Auth(), handlers.BookIndex)
	book.POST("/", middlewares.Auth(), handlers.BookStore)
	book.GET("/:id", middlewares.Auth(), handlers.BookFind)
	book.PUT("/:id", middlewares.Auth(), handlers.BookUpdate)
	book.DELETE("/:id", middlewares.Auth(), handlers.BookDestroy)
}