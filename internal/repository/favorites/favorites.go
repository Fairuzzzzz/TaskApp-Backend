package favorites

import (
	"errors"

	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
	"github.com/Fairuzzzzz/taskapp/internal/models/recipes"
	"gorm.io/gorm"
)

func (r *repository) CheckFavorite(recipeID uint, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&favorites.Favorites{}).Where("recipe_id = ? AND user_id = ?", recipeID, userID).Count(&count).Error
	return count > 0, err
}

func (r *repository) CreateFavorite(favorites *favorites.Favorites) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var recipe recipes.Recipe
		if err := tx.First(&recipe, favorites.RecipeID).Error; err != nil {
			return errors.New("recipe not found")
		}

		exists, err := r.CheckFavorite(favorites.RecipeID, favorites.UserID)
		if err != nil {
			return err
		}

		if exists {
			return errors.New("recipe already favorited")
		}

		if err := tx.Create(favorites).Error; err != nil {
			return err
		}

		if err := tx.Model(&recipes.Recipe{}).Where("id = ?", favorites.RecipeID).Update("is_favorite", true).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) GetAll(userID uint) ([]favorites.Favorites, error) {
	var favorites []favorites.Favorites
	if err := r.db.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *repository) Delete(ID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var favorite favorites.Favorites
		if err := tx.First(&favorite, ID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&favorite).Error; err != nil {
			return err
		}

		var count int64
		if err := tx.Model(&favorites.Favorites{}).Where("recipe_id = ? AND id != ?", favorite.RecipeID, ID).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := tx.Model(&recipes.Recipe{}).Where("id = ?", favorite.RecipeID).Update("is_favorite", false).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *repository) DeleteByRecipeID(recipeID uint, userID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find the favorite first
		var favorite favorites.Favorites
		if err := tx.Where("recipe_id = ? AND user_id = ?", recipeID, userID).First(&favorite).Error; err != nil {
			return err
		}

		// Delete the favorite
		if err := tx.Delete(&favorite).Error; err != nil {
			return err
		}

		// Check if any other users have favorited this recipe
		var count int64
		if err := tx.Model(&favorites.Favorites{}).Where("recipe_id = ?", recipeID).Count(&count).Error; err != nil {
			return err
		}

		// If no more favorites exist, update recipe's is_favorite to false
		if count == 0 {
			if err := tx.Model(&recipes.Recipe{}).Where("id = ?", recipeID).Update("is_favorite", false).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
