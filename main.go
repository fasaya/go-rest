package main

import (
	"go-rest/database"
	"go-rest/router"
)

var (
	PORT = ":9090"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
