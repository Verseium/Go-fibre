package books

import (
	"github.com/gofiber/fiber/v2"
)

// data

type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
}

// Array

var books = []Book{
	{ID: "1", Title: "Book One"},
	{ID: "2", Title: "Book Two"},
}

func GetBooks(c *fiber.Ctx) error{
	//return c.SendString("All Books")

	return c.JSON(books)
	

}

func GetBook(c *fiber.Ctx) error{
	return c.SendString("Single Book")
}

func NewBook(c *fiber.Ctx) error {
	// Parse request body
	var newBook Book
	if err := c.BodyParser(&newBook); err != nil {
		return err
	}

	// Add new book to array
	books = append(books, newBook)

	// Return success response
	return c.SendString("New book added successfully")
}

func DeleteBook(c *fiber.Ctx) error{
	 return c.SendString("Delete Book")
}
