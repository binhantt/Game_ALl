package config

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"Game_ALL_go_bankend/middleware"
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

	// Thêm middleware giới hạn và log IP
	router.Use(middleware.LimitIPAccess())
	router.Use(middleware.LogIPToFile())

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
