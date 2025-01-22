package favorites

import (
	"net/http"
	"strconv"

	"github.com/Fairuzzzzz/taskapp/internal/models/favorites"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddToFavorites(c *gin.Context) {
	var request favorites.Favorites
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")
	request.UserID = userID

	if err := h.service.CreateFavorite(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetAllFavorites(c *gin.Context) {
	userID := c.GetUint("user_id")

	favorites, err := h.service.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func (h *Handler) DeleteFavorite(c *gin.Context) {
	favoriteID := c.Param("id")
	id, err := strconv.ParseUint(favoriteID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.DeleteFavorites(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "favorite deleted successfully",
	})
}

func (h *Handler) DeleteByRecipeID(c *gin.Context) {
	recipeID := c.Param("recipeId")
	id, err := strconv.ParseUint(recipeID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")

	if err := h.service.DeleteByRecipeID(uint(id), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "favorite deleted successfully",
	})
}
