package repositories

import (
	"book-api/middlewares"
	"book-api/structs"
	"database/sql"
)

func GetAllBook(dbParam *sql.DB) (result []structs.Book, err error) {
	var category structs.Category
	sql := "SELECT * FROM books ORDER BY id ASC"
	rows, err := dbParam.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
		if err != nil {
			panic(err)
		}
		category.ID = *book.CategoryID
		_ = GetCategoryById(dbParam, &category)
		book.Category = category
		result = append(result, book)
	}
	return
}

func InsertBook(dbParam *sql.DB, book *structs.Book) error {
	sql := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) Returning *`
	err := dbParam.QueryRow(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, middlewares.GetAuth()).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	var category structs.Category
	category.ID = *book.CategoryID
	_ = GetCategoryById(dbParam, &category)
	book.Category = category

	return err
}

func UpdateBook(dbParam *sql.DB, book *structs.Book) error {
	sql := `UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, modified_by = $9, modified_at = NOW() WHERE id = $10 RETURNING *`
	err := dbParam.QueryRow(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, middlewares.GetAuth(), book.ID).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	var category structs.Category
	category.ID = *book.CategoryID
	_ = GetCategoryById(dbParam, &category)
	book.Category = category

	return err
}

func DeleteBook(dbParam *sql.DB, book *structs.Book) error {
	sqlStatement := `DELETE FROM books WHERE id = $1`
	result, _ := dbParam.Exec(sqlStatement, book.ID)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func GetBookById(dbParam *sql.DB, book *structs.Book) error {
	sql := `SELECT * FROM books WHERE id = $1`
	err := dbParam.QueryRow(sql, book.ID).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if book.CategoryID != nil {
		var category structs.Category
		category.ID = *book.CategoryID
		_ = GetCategoryById(dbParam, &category)
		book.Category = category
	}

	return err
}