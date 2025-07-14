package models

import (
	"gorm.io/gorm"
	"time"
)

// ProductDetail represents a detail for a product (e.g., minimum requirements)
type ProductDetail struct {
	gorm.Model
	ProductID   uint      `gorm:"not null" json:"product_id"`
	Title       string    `gorm:"size:100;not null" json:"title"` // Ví dụ: Cấu hình tối thiểu
	Detail      string    `gorm:"size:1000" json:"detail"`
	Platform    string    `gorm:"size:50" json:"platform"`
	ReleaseDate time.Time `json:"release_date"`
}

// TableName specifies the table name for the ProductDetail model
func (ProductDetail) TableName() string {
	return "product_details"
} 