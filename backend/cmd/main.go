package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sekarasiewicz/nurse/backend/pkg/db"
	"github.com/sekarasiewicz/nurse/backend/pkg/handlers"
)

func main() {
	db.InitDatabase()

	router := gin.Default()

	handlers.SetRoutes(router)

	router.Run(":8080")
}
