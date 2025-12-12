package database

import (
	"best-portfolio-go/config"
	"best-portfolio-go/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	var err error

	DB, err = gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate tables
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.Project{},
		&models.ProjectImage{},
		&models.AboutCard{},
		&models.SkillCategory{},
		&models.ContactPageData{},
		&models.ContactMessage{},
	)
}

func GetDB() *gorm.DB {
	return DB
}

