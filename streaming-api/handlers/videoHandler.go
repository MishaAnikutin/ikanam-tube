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
	router.Use(CORSMiddleware())

	video_router := router.Group("/api/video")

	video_router.GET("/:name", videoHandler)

	return router
}

func videoHandler(c *gin.Context) {
	name := c.Param("name")

	Video := filepath.Join("static", "videos", name)

	log.Println(Video)

	if _, err := os.Stat(Video); err == nil {
		c.Header("Content-Type", "application/vnd.apple.mpegurl")
		c.File(Video)

	} else {
		c.Status(http.StatusNotFound)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
