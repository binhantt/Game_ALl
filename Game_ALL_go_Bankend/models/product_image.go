package models

import "gorm.io/gorm"

// ProductImage represents an image for a product
type ProductImage struct {
	gorm.Model
	ProductID uint   `gorm:"not null" json:"product_id"`
	ImageURL  string `gorm:"size:255;not null" json:"image_url"`
}

// TableName specifies the table name for the ProductImage model
func (ProductImage) TableName() string {
	return "product_images"
} 