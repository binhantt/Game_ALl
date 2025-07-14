package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController xử lý các yêu cầu liên quan đến người dùng
type UserController struct{}

// NewUserController tạo mới một instance của UserController
func NewUserController() *UserController {
	return &UserController{}
}

// GetAllUsers trả về danh sách tất cả người dùng
func (uc *UserController) GetAllUsers(c *gin.Context) {
	// TODO: Thay thế bằng logic lấy dữ liệu từ database
	users := []map[string]interface{}{
		{"id": 1, "username": "user1", "email": "user1@example.com"},
		{"id": 2, "username": "user2", "email": "user2@example.com"},
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

// GetUserByID trả về thông tin một người dùng theo ID
func (uc *UserController) GetUserByID(c *gin.Context) {
	// TODO: Lấy ID từ URL và xử lý logic
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get user by ID",
		"id":      c.Param("id"),
	})
}

// CreateUser tạo mới một người dùng
func (uc *UserController) CreateUser(c *gin.Context) {
	// TODO: Xử lý tạo mới người dùng
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created successfully",
	})
}
