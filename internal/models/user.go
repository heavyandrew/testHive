package models

import "time"

type User struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-" db:"password_hash"`
}

type UserAsset struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	AssetID   int       `json:"asset_id" db:"asset_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
