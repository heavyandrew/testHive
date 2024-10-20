package repository

import (
	"github.com/jmoiron/sqlx"
	"testHive/internal/models"
)

type AssetRepository struct {
	DB *sqlx.DB
}

func NewAssetRepository(db *sqlx.DB) *AssetRepository {
	return &AssetRepository{DB: db}
}

func (r *AssetRepository) CreateAsset(asset *models.Asset) (int, error) {
	query := `INSERT INTO assets (name, description, price) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, asset.Name, asset.Description, asset.Price).Scan(&asset.ID)
	if err != nil {
		return 0, err
	}
	return asset.ID, nil
}

func (r *AssetRepository) DeleteAsset(assetID int) error {
	query := `DELETE FROM assets WHERE id = $1`
	_, err := r.DB.Exec(query, assetID)
	return err
}

func (r *AssetRepository) FindAssets(search string) ([]models.Asset, error) {
	var assets []models.Asset
	var query string
	var args []interface{}

	if search == "" {
		query = `SELECT id, name, description, price FROM assets`
	} else {
		query = `SELECT id, name, description, price FROM assets WHERE name ILIKE $1 OR description ILIKE $1`
		args = append(args, "%"+search+"%")
	}

	if err := r.DB.Select(&assets, query, args...); err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *AssetRepository) AddUserAsset(userID, assetID int) error {
	query := `INSERT INTO user_assets (user_id, asset_id) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, userID, assetID)
	return err
}

func (r *AssetRepository) GetAssetByID(assetID int) (*models.Asset, error) {
	var asset models.Asset
	query := `SELECT id, name, description, price FROM assets WHERE id = $1`
	if err := r.DB.Get(&asset, query, assetID); err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *AssetRepository) GetUserAssets(userID int) ([]models.Asset, error) {
	var assets []models.Asset
	query := `SELECT a.id, a.name, a.description, a.price
              FROM assets a
              JOIN user_assets ua ON a.id = ua.asset_id
              WHERE ua.user_id = $1`
	if err := r.DB.Select(&assets, query, userID); err != nil {
		return nil, err
	}
	return assets, nil
}
