package rest

import (
	v1 "github.com/MishaAnikutin/metadata-api/internal/transport/rest/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handlers *v1.VideoHandlers) *gin.Engine {
	router := gin.Default()

	router.GET("/video/:id", handlers.GetVideoByID)
	router.GET("/video", handlers.GetVideoByTag)
	router.GET("/videos", handlers.GetAllVideos)

	return router
}
