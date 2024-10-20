package models

type Asset struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name" validate:"required"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price" validate:"gt=-1"`
}
