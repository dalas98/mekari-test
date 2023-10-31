package db

import (
	"github.com/dalas98/mekari-test/app/repositories"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) repositories.DBRepositoryInterface {
	return &DBRepository{db}
}

func NewPostgreSQLConn(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	err = goose.Up(sqlDB, "app/repositories/db/migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}
