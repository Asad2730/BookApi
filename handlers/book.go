package handlers

import (
	"strconv"

	"github.com/Asad2730/BookApi/interfaces"
	"github.com/Asad2730/BookApi/models"
	"github.com/gin-gonic/gin"
)

type BookHandlers struct {
	b interfaces.BookRepo
}

func NewBookHandler(b interfaces.BookRepo) *BookHandlers {
	return &BookHandlers{b: b}
}

func (h *BookHandlers) Create(c *gin.Context) {

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(402, err.Error())
		return
	}

	if err := h.b.Create(&book); err != nil {
		c.JSON(401, err.Error())
		return
	}

	c.JSON(201, "Book Created!")
}

func (h *BookHandlers) Read(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	book, err := h.b.Read(bookID)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	c.JSON(201, book)
}

func (h *BookHandlers) ReadAll(c *gin.Context) {

	books, err := h.b.ReadAll()
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	c.JSON(201, books)
}

func (h *BookHandlers) Delete(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := h.b.Delete(bookID); err != nil {
		c.JSON(404, err.Error())
		return
	}

	c.JSON(200, "Book Removed!")
}
