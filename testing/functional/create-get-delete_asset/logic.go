package create_get_delete_asset

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	funcional "testHive/testing/functional"
	"time"
)

type Asset struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type AssetResponse struct {
	ID int `json:"id"`
}

func GenerateRandomAsset() Asset {
	rand.Seed(time.Now().UnixNano())
	price := float64(rand.Intn(1000)) + rand.Float64()
	price = math.Round(price*100) / 100
	return Asset{
		Name:        fmt.Sprintf("Asset-%d", rand.Intn(10000)),
		Description: fmt.Sprintf("Description for asset %d", rand.Intn(10000)),
		Price:       price,
	}
}

func AddAsset(accessToken string, asset Asset) (int, error) {
	url := fmt.Sprintf("%s:/api/assets", funcional.Api_url)

	jsonData, err := json.Marshal(asset)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("while creating asset expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var assetResponse AssetResponse
	if err := json.NewDecoder(resp.Body).Decode(&assetResponse); err != nil {
		return 0, err
	}

	return assetResponse.ID, nil
}

func GetAssets(accessToken string) ([]Asset, error) {
	url := fmt.Sprintf("%s:/api/assets", funcional.Api_url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("while getting existing assets expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var assets []Asset
	if err := json.NewDecoder(resp.Body).Decode(&assets); err != nil {
		return nil, err
	}

	return assets, nil
}

func DeleteAsset(accessToken string, assetID int) error {
	url := fmt.Sprintf("%s:/api/assets/%d", funcional.Api_url, assetID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("while deleting asset expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	return nil
}
