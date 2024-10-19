package services

import (
	"testHive/internal/models"
	"testHive/internal/repository"
)

// AssetService служит для обработки операций с активами
type AssetService struct {
	AssetRepo *repository.AssetRepository
}

func NewAssetService(repo *repository.AssetRepository) *AssetService {
	return &AssetService{AssetRepo: repo}
}

func (s *AssetService) AddAsset(userID int, asset *models.Asset) error {
	asset.UserID = userID
	return s.AssetRepo.CreateAsset(asset)
}

func (s *AssetService) DeleteAsset(assetID int) error {
	return s.AssetRepo.DeleteAsset(assetID)
}

func (s *AssetService) GetUserAssets(userID int) ([]models.Asset, error) {
	return s.AssetRepo.GetAssetsByUserID(userID)
}
