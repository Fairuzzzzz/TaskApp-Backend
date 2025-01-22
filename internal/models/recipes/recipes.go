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
	Title       string `json:"Title"`
	Description string `json:"Description"`
	ImageURL    string `json:"ImageURL"`
	Ingredients string `json:"Ingredients"`
	Eattime     string `json:"Eattime"`
	From        string `json:"From"`
	Nutrition   string `json:"Nutrition"`
	Calories    int    `json:"Calories"`
	Rating      int    `json:"Rating"`
	IsFavorite  bool   `json:"IsFavorite"`
}

type RecipeRequestByID struct {
	ID uint `json:"id"`
}
