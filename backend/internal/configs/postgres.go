package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/wisaitas/todo-web/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		ENV.DB_HOST,
		ENV.DB_USER,
		ENV.DB_PASSWORD,
		ENV.DB_NAME,
		ENV.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := autoMigrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err := autoSeed(db); err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	log.Println("database connected successfully")
	return db
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Province{},
		&models.District{},
		&models.SubDistrict{},
		&models.Address{},
	); err != nil {
		return fmt.Errorf("error migrating database: %w", err)
	}

	log.Println("database migrated successfully")

	return nil
}

func autoSeed(db *gorm.DB) error {
	seedConfigs := []struct {
		model       interface{}
		filename    string
		destination interface{}
		entityName  string
	}{
		{&models.Province{}, ENV.PROVINCE_FILE_PATH, &[]models.Province{}, "provinces"},
		{&models.District{}, ENV.DISTRICT_FILE_PATH, &[]models.District{}, "districts"},
		{&models.SubDistrict{}, ENV.SUB_DISTRICT_FILE_PATH, &[]models.SubDistrict{}, "sub districts"},
	}

	for _, config := range seedConfigs {
		if err := seedIfEmpty(db, config.model, config.filename, config.destination, config.entityName); err != nil {
			return err
		}
	}

	log.Println("database seeded successfully")
	return nil
}

func seedIfEmpty(db *gorm.DB, model interface{}, filename string, destination interface{}, entityName string) error {
	var count int64
	if err := db.Model(model).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking %s: %w", entityName, err)
	}

	if count == 0 {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening %s file: %v", entityName, err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("error reading %s file: %v", entityName, err)
		}

		if err := json.Unmarshal(byteData, destination); err != nil {
			log.Fatalf("error unmarshaling %s: %v", entityName, err)
		}

		if err := db.CreateInBatches(destination, 100).Error; err != nil {
			return fmt.Errorf("error seeding %s: %w", entityName, err)
		}
	}

	return nil
}
