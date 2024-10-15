package repository

import (
	"github.com/jmoiron/sqlx"
	"myapp/internal/models"
)

// AssetRepository интерфейс для операций с активами
type AssetRepository struct {
	DB *sqlx.DB
}

func NewAssetRepository(db *sqlx.DB) *AssetRepository {
	return &AssetRepository{DB: db}
}

func (r *AssetRepository) CreateAsset(asset *models.Asset) error {
	query := `INSERT INTO assets (user_id, name, description, price) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.DB.QueryRow(query, asset.UserID, asset.Name, asset.Description, asset.Price).Scan(&asset.ID)
}

func (r *AssetRepository) DeleteAsset(assetID int) error {
	query := `DELETE FROM assets WHERE id = $1`
	_, err := r.DB.Exec(query, assetID)
	return err
}

func (r *AssetRepository) GetAssetsByUserID(userID int) ([]models.Asset, error) {
	var assets []models.Asset
	query := `SELECT id, user_id, name, description, price FROM assets WHERE user_id = $1`
	if err := r.DB.Select(&assets, query, userID); err != nil {
		return nil, err
	}
	return assets, nil
}
