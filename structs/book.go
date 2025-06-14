package structs

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
	ReleaseYear *int `json:"release_year"`
	Price *int `json:"price"`
	TotalPage *int `json:"total_page"`
	Thickness string `json:"thickness"`
	CategoryID *int `json:"category_id"`
	CreatedAt  interface{} `json:"created_at"`
	CreatedBy  interface{}    `json:"created_by"`
	ModifiedAt interface{} `json:"modified_at"`
	ModifiedBy interface{}    `json:"modified_by"`
	Category interface{}    `json:"category"`
}