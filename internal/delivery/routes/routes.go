package routes

import (
	"parkir/internal/delivery"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App           *gin.Engine
	OcrController *delivery.OCRHandler
}

func (c *RouteConfig) Setup() {
	c.SetupOcrRoute()
}

func (c *RouteConfig) SetupOcrRoute() {
	c.App.POST("/cek", c.OcrController.UploadOCR)
}
