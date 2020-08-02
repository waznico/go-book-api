package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/waznico/go-book-api/book"
	"github.com/waznico/go-book-api/internal/database"
)

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic(err.Error)
	}
	log.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	log.Println("Automigration completed")
}
