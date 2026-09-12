package main

import (
	"gotodo/database"
	"gotodo/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided, could not find environment variables")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	h := handlers.NewTaskHandler(db)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	tasks := router.Group("/tasks")
	{
		tasks.GET("", h.GetTasks)
		tasks.GET("/:id", h.GetTask)
		tasks.DELETE("/:id", h.DeleteTask)
		tasks.POST("", h.CreateTask)
		tasks.PUT("/:id", h.UpdateTask)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
