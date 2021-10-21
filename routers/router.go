package routers

import (
	"bracket-backend/routers/api"
	"github.com/gin-gonic/gin"
)

// InitRouter creates the main application API router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/song/:id", api.GetSongByID)
	return router
}
