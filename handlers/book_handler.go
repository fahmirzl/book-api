package handlers

import (
	"book-api/db"
	"book-api/repositories"
	"book-api/structs"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookIndex(c *gin.Context) {
	book, err := repositories.GetAllBook(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Book retrieved successfully",
		Error:   nil,
		Data:    book,
	})
}

func BookStore(c *gin.Context) {
	var book structs.Book

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	validations := map[string]string{}
	if book.Title == "" {
		validations["title"] = "The title field is required"
	}
	if book.Description == "" {
		validations["description"] = "The description field is required"
	}
	if book.ImageURL == "" {
		validations["image_url"] = "The image_url field is required"
	}
	if book.ReleaseYear == nil {
		validations["release_year"] = "The release_year field is required"
	} else if *book.ReleaseYear < 1980 || *book.ReleaseYear > 2024 {
		validations["release_year"] = "Release year must be between 1980 and 2024"
	}
	if book.Price == nil {
		validations["price"] = "The price field is required"
	}
	if book.TotalPage == nil {
		validations["total_page"] = "The total_page field is required"
	}
	if book.CategoryID == nil {
		validations["category_id"] = "The title category_id is required"
	}
	if len(validations) > 0 {
		c.JSON(http.StatusUnprocessableEntity, structs.Response{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    validations,
		})
		return
	}

	if *book.TotalPage > 100 {
		book.Thickness = "Tebal"
	} else {
		book.Thickness = "Tipis"
	}

	err = repositories.InsertBook(db.DB, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Book inserted successfully",
		Error:   nil,
		Data:    book,
	})
}

func BookFind(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repositories.GetBookById(db.DB, &book)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Book with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return;
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: fmt.Sprintf("Book with id %d successfully found", id),
		Error: nil,
		Data: book,
	})
}

func BookUpdate(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	book.ID = id

	validations := map[string]string{}
	if book.Title == "" {
		validations["title"] = "The title field is required"
	}
	if book.Description == "" {
		validations["description"] = "The description field is required"
	}
	if book.ImageURL == "" {
		validations["image_url"] = "The image_url field is required"
	}
	if book.ReleaseYear == nil {
		validations["release_year"] = "The release_year field is required"
	} else if *book.ReleaseYear < 1980 || *book.ReleaseYear > 2024 {
		validations["release_year"] = "Release year must be between 1980 and 2024"
	}
	if book.Price == nil {
		validations["price"] = "The price field is required"
	}
	if book.TotalPage == nil {
		validations["total_page"] = "The total_page field is required"
	}
	if book.CategoryID == nil {
		validations["category_id"] = "The title category_id is required"
	}
	if len(validations) > 0 {
		c.JSON(http.StatusUnprocessableEntity, structs.Response{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    validations,
		})
		return
	}

	if *book.TotalPage > 100 {
		book.Thickness = "Tebal"
	} else {
		book.Thickness = "Tipis"
	}

	err = repositories.UpdateBook(db.DB, &book)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Book with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response{
		Message: "Book updated successfully",
		Error:   nil,
		Data:    book,
	})
}

func BookDestroy(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repositories.DeleteBook(db.DB, &book)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, structs.Response{
			Message: fmt.Sprintf("Book with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Response{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Response {
		Message: "Book deleted successfully",
		Error: nil,
		Data: nil,
	})
}