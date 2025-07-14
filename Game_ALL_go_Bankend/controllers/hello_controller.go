package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

// NewHelloController tạo mới một instance của HelloController
func NewHelloController() *HelloController {
	return &HelloController{}
}

// SayHello trả về một lời chào đơn giản
func (hc *HelloController) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Xin chào! Đây là API của Game ALL Backend",
	})
}
