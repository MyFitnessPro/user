package user

import (
	"context"
	"log"
	"os"

	goFirebase "github.com/MyFitnessPro/firebase"
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

	// Read secrets file
	secretsFile, err := os.ReadFile("secrets.json")
	if err != nil {
		log.Fatalf("Failed to read secrets file: %v", err)
	}

	// Initialize Firebase client
	firebaseClient, err := goFirebase.NewFirebaseClient(ctx, "myfitnessprocoll", secretsFile)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase client: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

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
	r.GET("/user/get", func(c *gin.Context) { handleGetUserRequest(c, firebaseClient) })
	r.DELETE("/user/delete", func(c *gin.Context) { handleDeleteUserRequest(c, firebaseClient) })
	r.POST("/user/upsert", func(c *gin.Context) { handleUpsertUserRequest(c, firebaseClient) })

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	r.Run(":50001")
}
