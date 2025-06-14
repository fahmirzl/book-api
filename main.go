package main

import (
	"book-api/db"
	"book-api/routers"
	"os"
)

func main() {
	db.Init()
	defer db.DB.Close()

	routers.StartServer().Run(":" + os.Getenv("PORT"))
}