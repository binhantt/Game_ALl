package routes

import (
	"Game_ALL_go_bankend/config"
	"Game_ALL_go_bankend/controllers"
	"Game_ALL_go_bankend/routes"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(server *config.Server) {
	// Khởi tạo controllers
	helloController := controllers.NewHelloController()
	userController := controllers.NewUserController()

	// Public routes
	server.Router.GET("/api/hello", helloController.SayHello)

	// User routes
	userGroup := server.Router.Group("/api/users")
	{
		userGroup.GET("", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.POST("", userController.CreateUser)
	}

	// Đăng ký route cho admin
	routes.SetupAdminRoutes(server.Router)

	// Có thể thêm các route khác ở đây
}
