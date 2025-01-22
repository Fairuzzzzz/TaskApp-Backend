package favorites

import (
	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
	"github.com/gin-gonic/gin"
)

type service interface {
	CreateFavorite(favorites *favorites.Favorites) error
	GetAll(userID uint) ([]favorites.Favorites, error)
	DeleteFavorites(ID uint) error
	CheckFavorite(recipeID uint, userID uint) (bool, error)
	DeleteByRecipeID(recipeID uint, userID uint) error
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/favorites")
	route.POST("/add", h.AddToFavorites)
	route.GET("/all", h.GetAllFavorites)
	route.DELETE("/recipe/:recipeId", h.DeleteByRecipeID)
}
