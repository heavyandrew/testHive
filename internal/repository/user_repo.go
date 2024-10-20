package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"testHive/internal/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
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
