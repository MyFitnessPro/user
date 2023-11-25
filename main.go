package main

import (
	"context"
	"log"
	"os"

	goFirebase "github.com/MyFitnessPro/firebase"
	middleware "github.com/MyFitnessPro/middleware"
	_ "github.com/MyFitnessPro/user/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 	User Service API
// @version	1.0
// @description Api to manage user CRUD operations
// @host 	localhost:50001
// @BasePath /user/
func main() {
	ctx := context.Background()

	// Get secrets file path from environment variable
	secretsFilePath := os.Getenv("SECRETS_FILE_PATH")
	projectID := os.Getenv("PROJECT_ID")

	// Read secrets file
	secretsFile, err := os.ReadFile(secretsFilePath)
	if err != nil {
		log.Fatalf("Failed to read secrets file: %v", err)
	}

	// Initialize Firebase client
	firebaseClient, err := goFirebase.NewFirebaseClient(ctx, projectID, secretsFile)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase client: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()
	r.Use(middleware.ProcessRequestMiddleware(firebaseClient))

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Routes
	r.GET("/user/get", func(c *gin.Context) { HandleGetUserRequest(c, firebaseClient) })
	r.DELETE("/user/delete", func(c *gin.Context) { HandleDeleteUserRequest(c, firebaseClient) })
	r.POST("/user/upsert", func(c *gin.Context) { HandleUpsertUserRequest(c, firebaseClient) })

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	r.Run(":50001")
}
