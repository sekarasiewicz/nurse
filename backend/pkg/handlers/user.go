package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekarasiewicz/nurse/backend/pkg/db"
	"github.com/sekarasiewicz/nurse/backend/pkg/models"
)

func GetUsersHandler(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(200, users)
}

func PostUserHandler(c *gin.Context) {
	var payload models.UserPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Email: payload.Email, Name: payload.Name}
	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, user)
}
