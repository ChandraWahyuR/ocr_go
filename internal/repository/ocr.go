package repository

import (
	"context"
	"database/sql"
	"parkir/internal/entity"
	"parkir/utils"
	"time"

	"github.com/sirupsen/logrus"
)

type OCRRepo struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewRepoOcr(db *sql.DB, log *logrus.Logger) *OCRRepo {
	return &OCRRepo{
		db:  db,
		log: log,
	}
}

func (r *OCRRepo) SavedLicensePlate(ctx context.Context, ocr *entity.Ocr) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	query := `INSERT INTO plat_model(id, plat_nomor, is_exit, entered, exited) VALUES($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		ocr.ID,
		ocr.PlatNomor,
		ocr.IsExit,
		ocr.Entered,
		ocr.Exited,
	)

	if err != nil {
		return utils.ParsePQError(err)
	}

	return nil
}
