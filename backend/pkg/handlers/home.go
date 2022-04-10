package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekarasiewicz/nurse/backend/pkg/models"
)

func HomeHandler(c *gin.Context) {
	data := &models.User{ID: 1, Name: "test 2"}

	c.JSON(http.StatusOK, data)
}
