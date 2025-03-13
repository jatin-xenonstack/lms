package handlers

import (
	"library-management1/database"
	"library-management1/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param authLogin body models.AuthLogin true "Login credentials"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string "error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var authLogin models.AuthLogin

	if err := c.ShouldBind(&authLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userFound models.User
	database.DB.Where("email = ?", authLogin.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User Not Exist"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authLogin.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong Password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte("auth-api-jwt-secret"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
