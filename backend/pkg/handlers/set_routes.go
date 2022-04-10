package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func NoMethodHandler(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"code": http.StatusMethodNotAllowed, "message": "Method Not Allowed"})
}

func OptionsHandler(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func SetRoutes(router *gin.Engine) {
	router.SetTrustedProxies([]string{"0.0.0.0"})
	router.Use(OptionsHandler)

	router.NoRoute(NoMethodHandler)

	router.GET("/", HomeHandler)
	router.GET("/api/users", GetUsersHandler)
	router.POST("/api/users", PostUserHandler)
}
