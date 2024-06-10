package models

type CreateItemInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
