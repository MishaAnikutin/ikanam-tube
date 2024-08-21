package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	video_router := router.Group("/video")
	video_router.GET("/:name", videoHandler)

	return router
}

func videoHandler(c *gin.Context) {
	name := c.Param("name") + ".m3u8"

	Video := filepath.Join("..", "static", "videos", name)

	log.Println(Video)

	if _, err := os.Stat(Video); err == nil {
		c.Header("Content-Type", "application/vnd.apple.mpegurl")
		c.File(Video)

	} else {
		c.Status(http.StatusNotFound)
	}
}
