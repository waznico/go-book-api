package book

import (
	"log"
	"strconv"

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
	if id == "" {
		log.Printf("Bad request. No id was given.")
		ctx.SendStatus(400)
		return
	}

	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		ctx.SendStatus(404)
		return
	}
	ctx.JSON(book)
}

// NewBook adds a new book into database
func NewBook(ctx *fiber.Ctx) {
	var bookIn Book
	err := ctx.BodyParser(&bookIn)
	if err != nil {
		log.Printf("Bad request to insert a new book with title: %v, author: %v and rating: %v\n", bookIn.Title, bookIn.Author, strconv.Itoa(bookIn.Rating))
		ctx.SendStatus(400)
		return
	}
	db := database.DBConn
	db.Create(&bookIn)
	ctx.JSON(bookIn)
}

// UpdateBook updates the book of the given id if existent
func UpdateBook(ctx *fiber.Ctx) {
	paramID := ctx.Params("id")
	_, parsErr := strconv.Atoi(paramID)
	if (paramID == "") || (parsErr != nil) {
		log.Printf("Bad request. No or invalid id was given.")
		ctx.SendStatus(400)
		return
	}

	var bookUp Book
	err := ctx.BodyParser(&bookUp)
	if err != nil {
		log.Printf("Bad request to update the book with id: %v, title: %v, author: %v and rating: %v\n", paramID, bookUp.Title, bookUp.Author, strconv.Itoa(bookUp.Rating))
		ctx.SendStatus(400)
		return
	}

	var book Book
	db := database.DBConn
	db.Find(&book, paramID)
	if len(book.Title) == 0 {
		ctx.SendStatus(404)
		return
	}
	db.Model(&book).Update(&bookUp)
	ctx.JSON(bookUp)
}

// DeleteBook deletes book of the given id
func DeleteBook(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	if id == "" {
		log.Printf("Bad request. No id was given.")
		ctx.SendStatus(400)
		return
	}

	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if len(book.Title) == 0 {
		ctx.SendStatus(404)
		return
	}

	db.Delete(&book, id)
	ctx.SendStatus(200)
}
