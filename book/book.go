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

// GetBooks godoc
// @Summary Get books stored in database
// @Description Get all books stored in database
// @Produce  json
// @Success 200 {object} []Book
// @Router /books [get]
func GetBooks(ctx *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	ctx.JSON(books)
}

// GetBook godoc
// @Summary Get book with id [id]
// @Description Returns single book with given id in database if it exists
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Router /book/{id} [get]
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

// NewBook godoc
// @Summary Adds a new book
// @Description Adds a new book into database
// @Accept  json
// @Produce	json
// @Param book body Book true "Book object to insert into"
// @Success 200 {object} Book
// @Failure 400 "Bad request"
// @Router /book [post]
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

// UpdateBook godoc
// @Summary Updates existing book
// @Description Updates the book of the given id if it exists. Otherwise an error will be thrown.
// @Accept  json
// @Produce	json
// @Param book body Book true "Book object with updated props (could also only an object containing the updated props)"
// @Success 200 {object} Book "Updated book"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Router /book/{id} [put]
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

// DeleteBook godoc
// @Summary Deletes a book
// @Description Deletes book with the given id if it exists in database.
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 "OK"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Router /book/{id} [delete]
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
