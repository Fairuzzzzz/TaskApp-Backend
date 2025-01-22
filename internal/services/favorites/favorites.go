package favorites

import (
	"errors"

	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateFavorite(favorites *favorites.Favorites) error {
	if favorites.RecipeID == 0 || favorites.UserID == 0 {
		return errors.New("recipe id and user id is required")
	}

	exists, err := s.CheckFavorite(favorites.RecipeID, favorites.UserID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("recipe already favorited")
	}

	return s.repository.CreateFavorite(favorites)
}

func (s *service) CheckFavorite(recipeID uint, userID uint) (bool, error) {
	if recipeID == 0 || userID == 0 {
		return false, errors.New("recipe id and user id is required")
	}

	return s.repository.CheckFavorite(recipeID, userID)
}

func (s *service) GetAll(userID uint) ([]favorites.Favorites, error) {
	if userID == 0 {
		return nil, errors.New("invalid user id")
	}

	favorites, err := s.repository.GetAll(userID)
	if err != nil {
		log.Error().Err(err).Msg("error getting favorites from database")
		return nil, err
	}
	return favorites, nil
}

func (s *service) DeleteFavorites(ID uint) error {
	if ID == 0 {
		return errors.New("invalid favorites id")
	}

	err := s.repository.Delete(ID)
	if err != nil {
		log.Error().Err(err).Msg("error deleting favorites from database")
		return err
	}
	return nil
}

func (s *service) DeleteByRecipeID(recipeID uint, userID uint) error {
	if recipeID == 0 || userID == 0 {
		return errors.New("recipe id and user id is required")
	}

	err := s.repository.DeleteByRecipeID(recipeID, userID)
	if err != nil {
		log.Error().Err(err).Msg("error deleting favorites from database")
		return err
	}
	return nil
}
