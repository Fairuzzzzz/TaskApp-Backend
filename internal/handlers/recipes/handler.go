package recipes

import (
	recipesModel "github.com/Fairuzzzzz/taskapp/internal/models/recipes"
	"github.com/gin-gonic/gin"
)

type service interface {
	CreateRecipes(request recipesModel.Recipe) error
	GetAll() ([]recipesModel.RecipeResponse, error)
	GetByID(id uint) (*recipesModel.RecipeResponse, error)
	UpdateRecipes(request recipesModel.Recipe, id uint) error
	DeleteRecipes(request recipesModel.RecipeRequestByID) error
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
	route := h.Group("/recipes")
	route.POST("/add", )
}
