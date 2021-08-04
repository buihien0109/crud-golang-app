package book

import (
	"fmt"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

var DB *pg.DB

func ConnectDatabase() *pg.DB {
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5434",
		User:     "postgres",
		Password: "123",
		Database: "postgres",
	})
	return DB
}

type Book struct {
	tableName struct{} `pg:"book.book"`
    Id int `pg:"type:serial,pk"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}


func GetBooks(c *fiber.Ctx) error {
	var books []Book
	err := DB.Model(&books).Select()
    if err != nil {
        panic(err)
    }
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	var id = c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	var book Book
	err = DB.Model(&book).Where("id=?", idInt).Select()
	if err != nil {
        panic(err)
    }
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	var req Book
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	fmt.Println(req)

	book := &Book {
		Author: req.Author,
		Title: req.Title,
		Rating: req.Rating,
	}
	_, err := DB.Model(book).Insert()
	if err != nil {
        panic(err)
    }

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	book := new(Book)
	_, err = DB.Model(book).Where("id = ?", idInt).Delete()
	if err != nil {
		panic(err)
	}
	return c.JSON("Book Successfully deleted")
}