package favorites

import (
	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
)

func (r *repository) CreateFavorite(favorites *favorites.Favorites) error {
	return r.db.Create(&favorites).Error
}

func (r *repository) GetAll(userID uint) ([]favorites.Favorites, error) {
	var favorites []favorites.Favorites
	if err := r.db.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *repository) Delete(ID uint) error {
	return r.db.Delete(&favorites.Favorites{}, ID).Error
}
