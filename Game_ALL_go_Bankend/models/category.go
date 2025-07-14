package models

import "gorm.io/gorm"

// Category dùng để phân loại sản phẩm/game
type Category struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products"`
}

func (Category) TableName() string {
	return "categories"
} 