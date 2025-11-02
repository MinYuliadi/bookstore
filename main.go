package main

import (
	"bookstore/config"
	"bookstore/routers"
)

func main() {
	config.InitDB()

	router := routers.SetupRouters()
	defer router.Run(":8080")
}
