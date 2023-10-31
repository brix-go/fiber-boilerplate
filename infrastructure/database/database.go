package database

import (
	"fmt"
	"github.com/brix-go/fiber/config"
	"github.com/brix-go/fiber/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func ConnectDatabase() (*gorm.DB, error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.AppConfig.Database.Host, config.AppConfig.Database.Port, config.AppConfig.Database.Username, config.AppConfig.Database.Password, config.AppConfig.Database.Database)
	dsn := fmt.Sprintf("postgresql://%s:%s@db:5432/%s?sslmode=disable", config.AppConfig.Database.Username, config.AppConfig.Database.Password, config.AppConfig.Database.Database)
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
	//	"localhost",
	//	config.AppConfig.Database.Port,
	//	config.AppConfig.Database.Username,
	//	config.AppConfig.Database.Password,
	//	config.AppConfig.Database.Database,
	//)
	fmt.Println("DSN : ", dsn)

	//define database logger
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dbConfig := gorm.Config{
		Logger: dbLogger,
	}
	//TODO : Change this driver if you use beside postgres (ex: mysql.Open())
	db, err := gorm.Open(postgres.Open(dsn), &dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		user.User{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
