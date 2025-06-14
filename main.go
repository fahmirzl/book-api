package main

import (
	"book-api/db"
	"book-api/routers"
)

func main() {
	db.Init()
	defer db.DB.Close()

	routers.StartServer().Run(":8080")
}