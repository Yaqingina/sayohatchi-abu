package database

import (
	"fmt"
	"go-telegram/config"
	"go-telegram/services/app/models"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	host := config.Config("POSTGRES_HOST")
	user := config.Config("POSTGRES_USER")
	pass := config.Config("POSTGRES_PASSWORD")
	port := config.Config("POSTGRES_PORT")
	dbName := config.Config("POSTGRES_DATABASE")
	ssl := config.Config("POSTGRES_SSL_MODE")

	defaultDSN := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=postgres sslmode=%s TimeZone=UTC",
		host, user, pass, port, ssl,
	)
	defaultDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to default DB: ", err)
	}

	var exists bool
	check := defaultDB.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", dbName).Scan(&exists)
	if check.Error != nil {
		log.Fatal("Failed to check db existence: ", check.Error)
	}

	if !exists {
		if err := defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error; err != nil {
			log.Fatal("Failed to create database: ", err)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=UTC",
		host, user, pass, port, dbName, ssl,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to target DB: ", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	return DB
}
