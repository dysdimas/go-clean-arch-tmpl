package db

import (
	"database/sql"
	"fmt"
	"github.com/dysdimas/go-clean-arch-tmpl/config"
	_ "github.com/go-sql-driver/mysql"
)

// GetDBConnection returns a database connection
func GetDBConnection() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
