package main

import (
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
        panic("failed to migrate database")
    }

    // Close the database connection
    dbSQL, err := initializers.DB.DB()
    if err != nil {
        panic("failed to close database connection")
    }
    defer dbSQL.Close()
}