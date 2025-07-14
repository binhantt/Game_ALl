package models

import (
	"gorm.io/gorm"
	"time"
)

// Order lưu thông tin lịch sử mua hàng của user
type Order struct {
	gorm.Model
	UserID     uint      `gorm:"not null" json:"user_id"`
	ProductID  uint      `gorm:"not null" json:"product_id"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	PurchasedAt time.Time `gorm:"autoCreateTime" json:"purchased_at"`
	User    User    `gorm:"foreignKey:UserID" json:"user"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

// TableName chỉ định tên bảng cho Order
func (Order) TableName() string {
	return "orders"
} 