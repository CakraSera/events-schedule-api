package routes

import (
	_ "example.com/rest-api/docs"
	"example.com/rest-api/middlewares"
	openapiui "github.com/PeterTakahashi/gin-openapi/openapiui"
	"github.com/gin-gonic/gin"
)

// Index godoc
// @Summary API Index
// @Description Welcome endpoint for the Events Schedule API
// @Tags general
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
// func index(context *gin.Context) {
// 	context.JSON(http.StatusOK, gin.H{
// 		"message": "Welcome to Events Schedule API",
// 		"version": "1.0",
// 		"docs":    "/docs",
// 	})
// }

func RegisterRouter(server *gin.Engine) {
	// Index route
	// server.GET("/", index)

	// Serve the OpenAPI JSON spec at /docs/openapi.json
	server.StaticFile("/docs/openapi.json", "./docs/swagger.json")

	// Serve API docs at root with Scalar UI
	server.GET("/", openapiui.WrapHandler(openapiui.Config{
		SpecURL:      "/docs/openapi.json",
		SpecFilePath: "./docs/swagger.json",
		Title:        "Events Schedule API",
		Theme:        "light", // or "dark"
	}))

	// Define a simple GET endpoint
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
