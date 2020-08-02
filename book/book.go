package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/waznico/go-book-api/internal/database"
)

// Book model struct
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks returns all books stored in database
func GetBooks(ctx *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	ctx.JSON(books)
}

// GetBook returns single book with given id in database if it exists
func GetBook(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		//TODO Error handling
	}
	ctx.JSON(book)
}

// NewBook adds a new book into database
func NewBook(ctx *fiber.Ctx) {
	ctx.Send("new book added")
}

// UpdateBook updates the book of the given id if existent
func UpdateBook(ctx *fiber.Ctx) {
	ctx.Send("book updated")
}

// DeleteBook deletes book of the given id
func DeleteBook(ctx *fiber.Ctx) {
	ctx.Send("book deleted")
}
