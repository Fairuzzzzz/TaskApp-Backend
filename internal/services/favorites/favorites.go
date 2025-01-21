package favorites

import (
	"errors"

	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateFavorite(favorites *favorites.Favorites) error {
	err := s.repository.CreateFavorite(favorites)
	if err != nil {
		log.Error().Err(err).Msg("error creating favorites in database")
		return err
	}
	return nil
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
