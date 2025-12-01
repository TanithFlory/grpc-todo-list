package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TanithFlory/go-todolist-api/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(appConfig *config.Config) (*sql.DB, error) {
	cfg := appConfig

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DATABASE,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	log.Println("MySQL connected")
	return db, nil
}
