package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	video_router := router.Group("/video")
	video_router.GET("/:name", getVideo)

	return router
}

func getVideo(c *gin.Context) {
	println("Отправка видео...")

	time.Sleep(10 * time.Second)

	println("Отправлено успешно!")
}
