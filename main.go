package main

import (
	"fmt"
	"log"
	"os"

	client "github.com/careyi3/crossfit_open_leaderboard/client"
	importer "github.com/careyi3/crossfit_open_leaderboard/importer"
	"github.com/careyi3/crossfit_open_leaderboard/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
		},
	)

	db, err := gorm.Open(sqlite.Open("data/crossfit_open_leaderboard.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Athlete{})
	db.AutoMigrate(&models.Result{})

	log.Println("Importing Men")
	for i := 1; i < 3390; i++ {
		response := client.FetchPage(2023, 1, i)
		importer.Import(response, db)
		msg := fmt.Sprintf("Imported Men Page: %d", i)
		log.Println(msg)
	}

	log.Println("Importing Women")
	for i := 1; i < 2657; i++ {
		response := client.FetchPage(2023, 2, i)
		importer.Import(response, db)
		msg := fmt.Sprintf("Imported Women Page: %d", i)
		log.Println(msg)
	}
}
