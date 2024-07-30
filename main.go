package main

import (
	"example.com/go-api/db"
	"example.com/go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8181")
}
