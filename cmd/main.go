package main

import (
	"blogpost/config"
	"blogpost/controllers"
	"blogpost/logger"
	"blogpost/routes"
	"blogpost/services"
	"log"
	"net/http"
)

func init() {
	config.LoadEnv()
}

func main() {
	// Connect to the database
	logger.Init()
	logger.Log.Info("Starting the application...")
	db := config.ConnectDB()

	// Initialize services
	articleService := services.NewArticleService(db)
	commentService := services.NewCommentService(db)

	// Initialize controllers
	articleController := controllers.NewArticleController(articleService)
	commentController := controllers.NewCommentController(commentService)

	// Setup routes
	router := routes.SetupRoutes(articleController, commentController)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
