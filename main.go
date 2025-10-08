package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Create a Gin router with default middleware (logger and recovery)
	server := gin.Default()

	routes.RegisterRouter(server)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	server.Run(":9090")
}
