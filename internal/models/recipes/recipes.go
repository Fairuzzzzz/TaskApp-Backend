package recipes

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	ImageURL    string `gorm:"not null"`
	Ingredients string `gorm:"not null"`
	Eattime     string `gorm:"not null"`
	From        string `gorm:"not null"`
	Nutrition   string `gorm:"not null"`
	Calories    int    `gorm:"not null"`
	Rating      int    `gorm:"not null"`
	IsFavorite  bool   `gorm:"default:false"`
}

type RecipeResponse struct {
	ID          uint   `json:"id"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	ImageURL    string `gorm:"not null"`
	Ingredients string `gorm:"not null"`
	Eattime     string `gorm:"not null"`
	From        string `gorm:"not null"`
	Nutrition   string `gorm:"not null"`
	Calories    int    `gorm:"not null"`
	Rating      int    `gorm:"not null"`
	IsFavorite  bool   `gorm:"default:false"`
}

type RecipeRequestByID struct {
	ID uint `json:"id"`
}
