package favorites

import "gorm.io/gorm"

type Favorites struct {
	gorm.Model
	RecipeID uint `gorm:"not null"`
}