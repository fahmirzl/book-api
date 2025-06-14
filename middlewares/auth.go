package middlewares

import (
	"book-api/db"
	"book-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	auth string
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, _ := c.Request.BasicAuth()

		var dbUsername, dbPassword string
		err := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", username).
			Scan(&dbUsername, &dbPassword)

		if err != nil || password != dbPassword {
			c.JSON(http.StatusUnauthorized, structs.Response{
				Message: "Invalid username or password",
				Error: "Unauthorized",
				Data: nil,
			})
			c.Abort()
			return
		}

		auth = dbUsername
		c.Next()
	}
}

func GetAuth() interface{} {
	return auth
}