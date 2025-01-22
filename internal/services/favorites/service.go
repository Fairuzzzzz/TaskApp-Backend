package favorites

import (
	"github.com/Fairuzzzzz/taskapp/internal/configs"
	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
)

type repository interface {
	CreateFavorite(favorites *favorites.Favorites) error
	GetAll(userID uint) ([]favorites.Favorites, error)
	Delete(ID uint) error
	CheckFavorite(recipeID uint, userID uint) (bool, error)
	DeleteByRecipeID(recipeID uint, userID uint) error
}

type service struct {
	cfg        *configs.Config
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg:        cfg,
		repository: repository,
	}
}
