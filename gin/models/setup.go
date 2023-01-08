package models

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	DSN := "host=localhost user=postgres password=postgres dbname=todo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Conneciton error: ", err)
	}

	DB.Migrator().DropTable(&Todo{})

	postgresDB, _ := DB.DB()

	postgresDB.SetMaxIdleConns(25)
	postgresDB.SetMaxOpenConns(100)
	postgresDB.SetConnMaxLifetime(time.Minute)

	DB.AutoMigrate(&Todo{})
}
