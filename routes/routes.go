package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) {
	// Define a simple GET endpoint
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
