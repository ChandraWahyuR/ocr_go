package migrations

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"
)

func CreateTable(db *sql.DB) error {
	files := []string{
		"./db/migrations/001_PlatModel.sql",
		"./db/migrations/002_Denda.sql",
	}

	for _, valuesFile := range files {
		query, err := os.ReadFile(valuesFile)
		if err != nil {
			log.Printf("Gagal membaca file %s: %s", valuesFile, err)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := db.ExecContext(ctx, string(query))
		if err != nil {
			log.Printf("Error %s when creating table", err)
			continue
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when getting rows affected", err)
			continue
		}
		log.Printf("Rows affected when creating table: %d", rows)
	}
	return nil
}
