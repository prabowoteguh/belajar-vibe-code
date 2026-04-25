package database

import (
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/prabowoteguh/belajar-vibe-code/config"
)

func InitSQLServer(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)

	db, err := sqlx.Connect("sqlserver", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
