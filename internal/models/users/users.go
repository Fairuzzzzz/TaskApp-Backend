package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
}
