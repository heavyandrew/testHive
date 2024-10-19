package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testHive/internal/models"
)

// UserRepository интерфейс для операций с пользователями
type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	_, err := r.GetUserByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("user already exists")
	}
	return r.DB.QueryRow(query, user.Username, user.PasswordHash).Scan(&user.ID)
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password_hash FROM users WHERE username = $1`
	if err := r.DB.Get(&user, query, username); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
