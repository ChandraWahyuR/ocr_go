package delivery

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OCRUsecaseInterface interface {
	SavedLicensePlate(ctx context.Context, image multipart.File, fileName string) error
	OCRFromForm(ctx context.Context, file multipart.File, fileName string) (string, float64, error)
}

type OCRHandlerInterface interface {
	SavedLicensePlate(c *gin.Context)
	OCRFromForm(c *gin.Context)
}

type OCRHandler struct {
	s OCRUsecaseInterface
}

func NewOcrHandler(s OCRUsecaseInterface) *OCRHandler {
	return &OCRHandler{
		s: s,
	}
}

func (h *OCRHandler) UploadOCR(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer file.Close()

	text, acc, err := h.s.OCRFromForm(c.Request.Context(), file, fileHeader.Filename)
	if err != nil {
		log.Println("OCR error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process OCR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"plat":     text,
		"accuracy": acc,
	})
}
