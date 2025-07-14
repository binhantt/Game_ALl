package models

import "gorm.io/gorm"

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null;unique" json:"username"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	FullName string `gorm:"size:100" json:"full_name"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
