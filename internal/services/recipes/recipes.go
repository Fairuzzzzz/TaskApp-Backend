package recipes

import (
	"errors"

	recipesModel "github.com/Fairuzzzzz/taskapp/internal/models/recipes"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) CreateRecipes(request recipesModel.Recipe) error {
	if request.Title == "" {
		return errors.New("title is required")
	}

	model := recipesModel.Recipe{
		Title:       request.Title,
		Description: request.Description,
		Ingredients: request.Ingredients,
		Eattime:     request.Eattime,
		From:        request.From,
		Nutrition:   request.Nutrition,
		Calories:    request.Calories,
		Rating:      request.Rating,
	}

	return s.repository.CreateRecipes(model)
}

func (s *service) GetAll() ([]recipesModel.RecipeResponse, error) {
	recipes, err := s.repository.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("error getting recipes from database")
		return nil, err
	}
	response := make([]recipesModel.RecipeResponse, len(recipes))
	for i, recipe := range recipes {
		response[i] = recipesModel.RecipeResponse{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Description: recipe.Description,
			ImageURL:    recipe.ImageURL,
			Ingredients: recipe.Ingredients,
			Eattime:     recipe.Eattime,
			From:        recipe.From,
			Nutrition:   recipe.Nutrition,
			Calories:    recipe.Calories,
			Rating:      recipe.Rating,
		}
	}

	return response, nil
}

func (s *service) GetByID(id uint) (*recipesModel.RecipeResponse, error) {
	if id == 0 {
		return nil, errors.New("invalid recipe id")
	}

	request, err := s.repository.GetByID(id)
	if err != nil {
		log.Error().Err(err).Msg("error getting recipe from database")
		return nil, err
	}

	response := &recipesModel.RecipeResponse{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		ImageURL:    request.ImageURL,
		Ingredients: request.Ingredients,
		Eattime:     request.Eattime,
		From:        request.From,
		Nutrition:   request.Nutrition,
		Calories:    request.Calories,
		Rating:      request.Rating,
	}
	return response, nil
}

func (s *service) UpdateRecipes(request recipesModel.Recipe, id uint) error {
	if request.Title == "" {
		return errors.New("title is required")
	}

	existingRecipes, err := s.repository.GetByID(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error getting recipes from database")
		return err
	}

	model := recipesModel.Recipe{
		Model: gorm.Model{
			ID: existingRecipes.ID,
		},
		Title: request.Title,

		Description: func() string {
			if request.Description != "" {
				return request.Description
			}
			return existingRecipes.Description
		}(),

		ImageURL: func() string {
			if request.ImageURL != "" {
				return request.ImageURL
			}
			return existingRecipes.ImageURL
		}(),

		Ingredients: func() string {
			if request.Ingredients != "" {
				return request.Ingredients
			}
			return existingRecipes.Ingredients
		}(),

		Eattime: func() string {
			if request.Eattime != "" {
				return request.Eattime
			}
			return existingRecipes.Eattime
		}(),

		From: func() string {
			if request.From != "" {
				return request.From
			}
			return existingRecipes.From
		}(),

		Nutrition: func() string {
			if request.Nutrition != "" {
				return request.Nutrition
			}
			return existingRecipes.Nutrition
		}(),

		Calories: func() int {
			if request.Calories > 0 {
				return request.Calories
			}
			return existingRecipes.Calories
		}(),

		Rating: func() int {
			if request.Rating > 0 && request.Rating <= 5 {
				return request.Rating
			}
			return existingRecipes.Rating
		}(),
	}

	if err := s.repository.Update(model); err != nil {
		log.Error().Err(err).Msg("error updating recieps in database")
		return err
	}
	return nil
}

func (s *service) DeleteRecipes(request recipesModel.RecipeRequestByID) error {
	err := s.repository.Delete(request.ID)
	if err != nil {
		log.Error().Err(err).Msg("error delete recipe from database")
		return err
	}
	return nil
}
