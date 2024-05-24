package db

import (
	"database/sql"
	"fmt"
	"os"
	"real_estate/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	sqlDB, err := sql.Open("postgres", config.Conf.DatabaseURL)
	if err != nil {
		fmt.Println("Unable to open postgres connection. Err:", err)
		os.Exit(1)
	}

	sqlDB.SetConnMaxIdleTime(5)
	sqlDB.SetMaxOpenConns(config.Conf.MaxDBConn)
	sqlDB.SetConnMaxLifetime(10)

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to create Gorm connection. Err:", err)
	}
}

type DB struct {
	*gorm.DB
}

func New() *DB {
	return &DB{
		DB: db,
	}
}
