package recipes

import (
	"net/http"
	"strconv"

	"github.com/Fairuzzzzz/taskapp/internal/models/recipes"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddRecipes(c *gin.Context) {
	var request recipes.Recipe
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.CreateRecipes(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) UpdateRecipes(c *gin.Context) {
	recipesID := c.Param("id")
	id, err := strconv.ParseUint(recipesID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request recipes.Recipe
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.UpdateRecipes(request, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
