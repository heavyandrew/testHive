package models

// User представляет пользователя в системе
type User struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Password     string `json:"password,omitempty"`   // не хранить пароль в базе данных
	PasswordHash string `json:"-" db:"password_hash"` // хэш пароля
}
