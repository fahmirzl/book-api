package structs

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt  interface{} `json:"created_at"`
	CreatedBy  interface{}    `json:"created_by"`
	ModifiedAt interface{} `json:"modified_at"`
	ModifiedBy interface{}    `json:"modified_by"`
}