package main

import (
	"helthmed-appointment/internal/appointment/infrastructure/db"
	"helthmed-appointment/internal/appointment/infrastructure/rabbitmq"
	"helthmed-appointment/internal/appointment/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Initialize RabbitMQ connection
	rabbitmq.Init()
	defer rabbitmq.Close()

	// Set up Gin router
	router := gin.Default()

	// Initialize HTTP handlers and routes
	http.InitRoutes(router)

	// Run the server
	router.Run(":8082")
}
