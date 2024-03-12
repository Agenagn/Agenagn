package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	log.Println("Connecting to database", os.Getenv("DB_URL"))
	var err error
	dsn := "postgresql://minasesotlg:9LJohpN4BeZH@ep-square-snow-a2yuweuk.eu-central-1.aws.neon.tech/Agenagn?sslmode=require"
	log.Println("Connecting to database")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	log.Println("Connected to database")
}