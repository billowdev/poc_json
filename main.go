package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		DB_HOST,
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)
	ConDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	if DB_RUN_MIGRATION {
		if err := RunMigrations(ConDB); err != nil {
			panic("failed to connect database")
		}
	}
	if DB_RUN_SEEDER {
		RunSeeds(ConDB)
	}

	repo := NewRepository(ConDB)
	srv := NewService(repo)
	handler := NewHandler(srv)
	_ = handler
}
