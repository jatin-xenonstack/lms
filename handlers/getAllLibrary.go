package handlers

import (
	"library-management1/database"
	"library-management1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllLibrary godoc
// @Summary Get all libraries
// @Description Get details of all libraries
// @Tags libraries
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Library
// @Router /getalllibrary [get]
func GetAllLibrary(c *gin.Context) {
	var alllibrary []models.Library

	database.DB.Find(&alllibrary)

	c.JSON(http.StatusOK, gin.H{"data": alllibrary})
}
