package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"parkir/internal/entity"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type OCRRepoInterface interface {
	SavedLicensePlate(ctx context.Context, ocr *entity.Ocr) error
}

type OCRProcessor interface {
	UploadFile(ctx context.Context, image multipart.File, fileName string) (string, float64, error)
}

type OcrUsecase struct {
	ocr  OCRProcessor
	repo OCRRepoInterface
	log  *logrus.Logger
}

func NewUseCaseOcr(repo OCRRepoInterface, log *logrus.Logger, ocr OCRProcessor) *OcrUsecase {
	return &OcrUsecase{
		repo: repo,
		log:  log,
		ocr:  ocr,
	}
}

func (s *OcrUsecase) SavedLicensePlate(ctx context.Context, image multipart.File, fileName string) error {
	text, acc, err := s.ocr.UploadFile(ctx, image, fileName)
	if err != nil {
		return err
	}

	fmt.Println(acc)
	stored := &entity.Ocr{
		ID:        uuid.New().String(),
		PlatNomor: text,
		IsExit:    false,
		Entered:   time.Now(),
	}

	go func() {
		_ = s.repo.SavedLicensePlate(context.Background(), stored)
	}()

	return nil
}

func (s *OcrUsecase) OCRFromForm(ctx context.Context, file multipart.File, fileName string) (string, float64, error) {
	text, acc, err := s.ocr.UploadFile(ctx, file, fileName)
	if err != nil {
		return "", 0, err
	}

	fmt.Println("OCR Text:", text)
	fmt.Println("Accuracy:", acc)
	return text, acc, nil
}
