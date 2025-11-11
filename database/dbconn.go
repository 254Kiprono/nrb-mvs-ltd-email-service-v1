package database

import (
	"email-service/config"
	"email-service/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitializeDB sets up the database connection
func InitializeDB(cfg *config.Config) (*gorm.DB, error) {
	// Build DSN with optimized parameters
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC&timeout=5s&readTimeout=30s&writeTimeout=30s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Configure GORM with production settings
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		SkipDefaultTransaction: true, // Disable implicit transactions for migrations
	}

	// Establish connection with retry
	var db *gorm.DB
	var err error
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
		if err == nil {
			break
		}
		log.Printf("Database connection attempt %d failed: %v", i+1, err)
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	// Run migrations
	if err := MigrateModels(db); err != nil {
		return nil, fmt.Errorf("database migration failed: %w", err)
	}

	DB = db
	return db, nil
}

// MigrateModels runs the database migrations for all models
func MigrateModels(db *gorm.DB) error {
	// Disable foreign key checks for faster migrations
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		return err
	}
	defer db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// Create tables with optimized schema
	models := []interface{}{
		&models.ContactMessage{},
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, model := range models {
			if err := tx.AutoMigrate(model); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
