package routes

import (
	"Game_ALL_go_bankend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("/api/admin")
	{
		adminGroup.POST("/send-verification", controllers.SendVerificationCodeHandler)
		adminGroup.POST("/verify-code", controllers.VerifyCodeHandler)
	}
}
