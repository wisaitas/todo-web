package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/wisaitas/todo-web/internal/models"
	"golang.org/x/crypto/bcrypt"

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
		{&models.Role{}, ENV.ROLES_FILE_PATH, &[]models.Role{}, "roles"},
	}

	tx := db.Begin()

	for _, config := range seedConfigs {
		if err := seedIfEmpty(tx, config.model, config.filename, config.destination, config.entityName); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := seedUsers(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error committing transaction: %w", err)
	}

	log.Println("database seeded successfully")
	return nil
}

func seedIfEmpty(tx *gorm.DB, model interface{}, filename string, destination interface{}, entityName string) error {
	var count int64
	if err := tx.Model(model).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking %s: %w", entityName, err)
	}

	if count == 0 {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("error opening %s file: %w", entityName, err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("error reading %s file: %w", entityName, err)
		}

		if err := json.Unmarshal(byteData, destination); err != nil {
			return fmt.Errorf("error unmarshaling %s: %w", entityName, err)
		}

		if err := tx.CreateInBatches(destination, 100).Error; err != nil {
			return fmt.Errorf("error seeding %s: %w", entityName, err)
		}
	}

	return nil
}

func seedUsers(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&models.User{}).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking %s: %w", "users", err)
	}

	if count > 0 {
		return nil
	}

	file, err := os.Open(ENV.USERS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error opening %s file: %w", "users", err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading %s file: %w", "users", err)
	}

	var users []models.User

	if err := json.Unmarshal(byteData, &users); err != nil {
		return fmt.Errorf("error unmarshaling %s: %w", "users", err)
	}

	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}

		var role models.Role
		if err := tx.First(&role, "name = ?", user.Role.Name).Error; err != nil {
			return fmt.Errorf("error finding role: %w", err)
		}

		user.Role = &role
		user.Password = string(hashedPassword)

		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("error seeding %s: %w", "users", err)
		}
	}

	return nil
}
