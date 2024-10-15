package models

// Asset представляет актив, принадлежащий пользователю
type Asset struct {
	ID          int     `json:"id" db:"id"`
	UserID      int     `json:"user_id" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
}
