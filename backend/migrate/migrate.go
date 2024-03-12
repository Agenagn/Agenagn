package main

import (
	"fmt"

	"example.com/agenagn/initializers"
	"example.com/agenagn/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Company{},
		&models.JobSeeker{},
		&models.Job{},
		&models.Application{},
		&models.EducationInfo{},
		&models.Bookmark{},
		&models.Sector{},
		&models.JobSector{},
	)

	if err != nil {
		fmt.Println("Failed to migrate database:", err)
		return
	}

	fmt.Println("Database migration successful.")

	// Close the database connection
	dbSQL, err := initializers.DB.DB()
	if err != nil {
		fmt.Println("Failed to close database connection:", err)
		return
	}
	defer dbSQL.Close()
}
