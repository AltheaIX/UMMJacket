package infras

import (
	"fmt"
	"github.com/AltheaIX/UMMJacket/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func InitPostgres(cfg *configs.Config) (*sqlx.DB, error) {
	var db *sqlx.DB

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
