package database

import (
	"fmt"
	"log"
	"testHive/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Connect подключается к базе данных PostgreSQL, используя параметры из config.Config
func Connect(cfg *config.Config) (*sqlx.DB, error) {
	// Формирование строки подключения из параметров конфигурации
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	// Подключение к базе данных с использованием сформированной строки подключения
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Проверка доступности подключения к базе данных
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к базе данных")
	return db, nil
}
