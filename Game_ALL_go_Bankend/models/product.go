package models

import (
	"gorm.io/gorm"
)

// Product represents a game or product in the system
type Product struct {
	gorm.Model
	Name        string         `gorm:"size:100;not null;unique" json:"name"`
	Price       float64        `gorm:"not null" json:"price"`
	Description string         `gorm:"size:500" json:"description"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
	Images      []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
	Details     []ProductDetail `gorm:"foreignKey:ProductID" json:"details"`
}

// TableName specifies the table name for the Product model
func (Product) TableName() string {
	return "products"
} 