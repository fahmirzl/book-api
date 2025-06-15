package repositories

import (
	"book-api/middlewares"
	"book-api/structs"
	"database/sql"
)

func GetAllCategory(dbParam *sql.DB) (result []structs.Category, err error) {
	sql := `SELECT * FROM categories ORDER BY id ASC`
	rows, err := dbParam.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
		if err != nil {
			panic(err)
		}
		result = append(result, category)
	}
	return
}

func InsertCategory(dbParam *sql.DB, category *structs.Category) error {
	sql := `INSERT INTO categories(name, created_by) VALUES($1, $2) Returning *`
	err := dbParam.QueryRow(sql, category.Name, middlewares.GetAuth()).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	return err
}

func UpdateCategory(dbParam *sql.DB, category *structs.Category) error {
	sql := `UPDATE categories SET name = $1, modified_at = NOW(), modified_by = $2 WHERE id = $3 Returning *`
	err := dbParam.QueryRow(sql, category.Name, middlewares.GetAuth(), category.ID).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	return err
}

func DeleteCategory(dbParam *sql.DB, category *structs.Category) error {
	sqlStatement := `DELETE FROM categories WHERE id = $1`
	result, _ := dbParam.Exec(sqlStatement, category.ID)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func GetCategoryById(dbParam *sql.DB, category *structs.Category) error {
	sql := `SELECT * FROM categories WHERE id = $1`
	err := dbParam.QueryRow(sql, category.ID).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	return err
}

func GetBookByCategoryId(dbParam *sql.DB, category *structs.Category) (result []structs.Book, err error) {
	sql := `SELECT books.* FROM books INNER JOIN categories ON books.category_id = categories.id where categories.id = $1`
	rows, err := dbParam.Query(sql, category.ID)
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
		result = append(result, book)
	}
	return
}
