package main

import (
	"github.com/abhilov23/gin_project/db"
	"github.com/abhilov23/gin_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	// here the default() is the gin router
	server := gin.Default()

	// this RegisterRoutes() handles all the routes
	routes.RegisterRoutes(server) 

	
	server.Run(":8080") // localhost 8080
}
