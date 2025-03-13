package handlers

import (
	"fmt"
	"library-management1/database"
	"library-management1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchBook godoc
// @Summary Search for books
// @Description Search for books by title, author, or publisher
// @Tags books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param search path string true "Search query"
// @Success 200 {object}  []models.BookInventory
// @Failure 400 {object} string
// @Router //book/{search} [get]
// @Security BearerAuth
func SearchBook(c *gin.Context) {
	var allBooks []models.BookInventory
	q := c.Param("search")
	fmt.Println(q)
	// search := c.Query(q)
	// fmt.Println(search)
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	database.DB.Model(&models.BookInventory{}).Where("title ILIKE ?", "%"+q+"%").
		Or("author ILIKE ?", "%"+q+"%").
		Or("publisher ILIKE ?", "%"+q+"%").
		Find(&allBooks)

	c.JSON(http.StatusOK, gin.H{"data": allBooks})
}
