package api

import (
	"bracket-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSongByID gets a song given an ID. Makes sense yeah?
func GetSongByID(c *gin.Context) {
	id, _ := c.Params.Get("id")
	song := service.GetSong(id)
	c.JSON(http.StatusOK, song)
}