package recipes

import (
	"github.com/Fairuzzzzz/taskapp/internal/configs"
	"github.com/Fairuzzzzz/taskapp/internal/models/recipes"
)

type repository interface {
	CreateRecipes(model recipes.Recipe) error
	GetAll() ([]recipes.Recipe, error)
	GetByID(id uint) (*recipes.Recipe, error)
	Update(model recipes.Recipe) error
	Delete(id uint) error
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
