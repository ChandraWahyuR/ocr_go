package app

import (
	"database/sql"
	"parkir/config"
	"parkir/internal/delivery"
	"parkir/internal/delivery/routes"
	"parkir/internal/ocr"
	"parkir/internal/repository"
	"parkir/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Boostrap struct {
	DB     *sql.DB
	App    *gin.Engine
	Log    *logrus.Logger
	Config *config.Config
}

func App(cfg *Boostrap) {
	grpcClient := config.InitGrpcOCRClient()
	ocrProcessor := ocr.NewOCRProcessor(grpcClient)

	ocrRepo := repository.NewRepoOcr(cfg.DB, cfg.Log)
	ocrUc := usecase.NewUseCaseOcr(ocrRepo, cfg.Log, ocrProcessor)
	ocrHdr := delivery.NewOcrHandler(ocrUc)

	routeConfig := routes.RouteConfig{
		App:           cfg.App,
		OcrController: ocrHdr,
	}

	routeConfig.Setup()
}
