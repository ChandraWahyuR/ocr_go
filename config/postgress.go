package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

func InitDatabase(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Database.dbUser, cfg.Database.dbPass, cfg.Database.dbHost, cfg.Database.dbPort, cfg.Database.dbName)
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", cfg.Database.dbUser, cfg.Database.dbPass, cfg.Database.dbHost, cfg.Database.dbPort, cfg.Database.dbName)
	conn, _ := url.Parse(dsn)
	// conn.RawQuery = fmt.Sprintf("sslmode=verify-ca&sslrootcert=%s", cfg.Database.cert)

	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfull connect to database on DSN:%s", dsn)
	return db, nil
}
