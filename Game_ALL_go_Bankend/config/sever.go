package config

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	Router *gin.Engine
	Port   string
}

// NewServer creates a new HTTP server
func NewServer(port string) *Server {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Add health check endpoint
	router.GET("/health", func(c *gin.Context) {
		status := "ok"
		dbStatus := "connected"
		if DB == nil {
			dbStatus = "disconnected"
		}

		c.JSON(200, gin.H{
			"status":    status,
			"message":   "Server is running",
			"database":  dbStatus,
			"timestamp": time.Now().Unix(),
		})
	})

	return &Server{
		Router: router,
		Port:   port,
	}
}

// Start runs the HTTP server
func (s *Server) Start() error {
	// Add your API routes here
	// Example: s.Router.GET("/api/endpoint", handlerFunction)

	// Start the server
	serverAddr := fmt.Sprintf(":%s", s.Port)
	log.Printf("Server is running on port %s\n", s.Port)
	return s.Router.Run(serverAddr)
}
