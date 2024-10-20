package services

import (
	"fmt"
	"testHive/internal/models"
	"testHive/internal/repository"
)

type AssetService struct {
	AssetRepo *repository.AssetRepository
}

func NewAssetService(repo *repository.AssetRepository) *AssetService {
	return &AssetService{AssetRepo: repo}
}

func (s *AssetService) AddAsset(asset *models.Asset) (int, error) {
	return s.AssetRepo.CreateAsset(asset)
}

func (s *AssetService) DeleteAsset(assetID int) error {
	return s.AssetRepo.DeleteAsset(assetID)
}

func (s *AssetService) FindAssets(search string) ([]models.Asset, error) {
	return s.AssetRepo.FindAssets(search)
}

func (s *AssetService) BuyAsset(userID, assetID int, price float64) error {
	asset, err := s.AssetRepo.GetAssetByID(assetID)
	if err != nil {
		return err
	}
	if asset.Price != price {
		return fmt.Errorf("цена актива не совпадает")
	}
	return s.AssetRepo.AddUserAsset(userID, assetID)
}

func (s *AssetService) GetUserAssets(userID int) ([]models.Asset, error) {
	return s.AssetRepo.GetUserAssets(userID)
}
