package main

import (
	"log"
	"net/http"

	"github.com/mukezhz/logging-template/handlers"
	"github.com/mukezhz/logging-template/models"
	"github.com/mukezhz/logging-template/repos"
	"github.com/mukezhz/logging-template/seeds"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("messages.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.MessageTemplate{})
	return db
}

func main() {
	db := InitDB()
	seeds.Seed(db)

	repo := repos.NewTemplateRepo(db)

	http.HandleFunc("GET /api/toast", handlers.ToastHandler(repo))

	http.ListenAndServe(":8080", nil)
}
