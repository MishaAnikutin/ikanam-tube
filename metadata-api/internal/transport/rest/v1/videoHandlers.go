package v1

import (
	"net/http"

	"github.com/MishaAnikutin/metadata-api/internal/service"
	"github.com/gin-gonic/gin"
)

type VideoHandlers struct {
	MetadataService *service.MetadataService
}

func New(m *service.MetadataService) *VideoHandlers {
	return &VideoHandlers{MetadataService: m}
}

func (h *VideoHandlers) GetVideoByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	metadata, err := h.MetadataService.GetByID(ctx, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if metadata == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metadata not found"})
		return
	}

	c.JSON(http.StatusOK, metadata)
}

func (h *VideoHandlers) GetVideoByTag(c *gin.Context) {
	ctx := c.Request.Context()
	tag := c.Query("tag")
	metadataList, err := h.MetadataService.GetByTag(ctx, tag)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if metadataList == nil || len(*metadataList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metadata not found"})
		return
	}

	c.JSON(http.StatusOK, metadataList)
}

func (h *VideoHandlers) GetAllVideos(c *gin.Context) {
	ctx := c.Request.Context()
	metadataList, err := h.MetadataService.GetAll(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if metadataList == nil || len(*metadataList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metadata not found"})
		return
	}

	c.JSON(http.StatusOK, metadataList)
}
