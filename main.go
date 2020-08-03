package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	//docpath for swagger doesn't need to be imported
	_ "github.com/waznico/go-book-api/docs"

	"github.com/waznico/go-book-api/book"
	"github.com/waznico/go-book-api/internal/database"
)

// @title Go Book API
// @version 1.0
// @description It's a sample API to demostrate how it could be created with go. You can manage books with it.
// @contact.name NiWA Dev
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1/
func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)
	app.Use("/swagger", swagger.Handler)

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
