package controllers

import (
	"net/http"
	"Game_ALL_go_bankend/services"
	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Email string `json:"email"`
}

type VerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// Gửi mã xác thực tới email
func SendVerificationCodeHandler(c *gin.Context) {
	var req EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	if err := services.SendVerificationCode(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send code"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent"})
}

// Xác thực mã code
func VerifyCodeHandler(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if services.VerifyCode(req.Email, req.Code) {
		c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired code"})
	}
}
