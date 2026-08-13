package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	handler_http "thekasir/internal/adapter/handler/http"
	"thekasir/internal/adapter/repository/postgres"
	"thekasir/internal/core/service"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	ctx := context.Background()

	// Initialize Database
	db, err := postgres.NewDB(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Pool.Close()

	// Auto-migrate (Simple Ponytail migration for MVP)
	log.Println("Running database migrations...")
	migrations := []string{"001_init.sql", "002_pos.sql"}
	for _, m := range migrations {
		migrationSQL, err := os.ReadFile("scripts/migrations/" + m)
		if err == nil {
			_, err = db.Pool.Exec(ctx, string(migrationSQL))
			if err != nil {
				log.Printf("Warning: Migration %s execution failed (could be already migrated): %v", m, err)
			} else {
				log.Printf("Migration %s executed successfully!", m)
			}
		} else {
			log.Printf("Warning: Could not read migration file %s: %v", m, err)
		}
	}

	// Initialize Repositories
	userRepo := postgres.NewUserRepository(db)
	_ = postgres.NewWorkspaceRepository(db)
	posRepo := postgres.NewPosRepository(db)
	analyticsRepo := postgres.NewAnalyticsRepository(db)

	// Initialize Services
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "supersecret_dev_key"
	}
	authService := service.NewAuthService(userRepo, jwtSecret)
	posService := service.NewPosService(posRepo)
	analyticsService := service.NewAnalyticsService(analyticsRepo)

	// Initialize Handlers
	authHandler := handler_http.NewAuthHandler(authService)
	posHandler := handler_http.NewPosHandler(posService)
	analyticsHandler := handler_http.NewAnalyticsHandler(analyticsService)

	// Setup Router
	r := gin.Default()

	// CORS Middleware could be added here

	// Routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "pong"})
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// POS Endpoints (Workspace Scoped)
		workspace := api.Group("/workspaces/:workspaceId")
		{
			workspace.GET("/products", posHandler.GetProducts)
			workspace.POST("/products", posHandler.CreateProduct)
			workspace.POST("/checkout", posHandler.Checkout)
			
			// Analytics
			workspace.GET("/analytics/summary", analyticsHandler.GetSummary)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on :%s", port)
	r.Run(":" + port)
}
