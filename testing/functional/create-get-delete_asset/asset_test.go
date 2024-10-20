package create_get_delete_asset

import (
	"log"
	"testHive/testing/functional/login"
	"testHive/testing/functional/register"
	"testing"
	"time"
)

func TestAddAndDeleteAsset(t *testing.T) {
	username, password := register.GenerateRandomUser()
	_, err := register.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("Error registering user: %v", err)
	}
	time.Sleep(2)

	accessToken, err := login.LoginUser(username, password)
	if err != nil {
		t.Fatalf("Error logging in: %v", err)
	}

	asset := GenerateRandomAsset()

	assetId, err := AddAsset(accessToken, asset)
	if err != nil {
		t.Fatalf("Error adding asset: %v", err)
	}

	time.Sleep(5)

	assets, err := GetAssets(accessToken)
	if err != nil {
		t.Fatalf("Error getting assets: %v", err)
	}

	found := false
	for _, a := range assets {
		log.Printf("Compare: %v, with %v", a, asset)
		if a.Name == asset.Name && a.Description == asset.Description && a.Price == asset.Price {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Expected to find asset %v, but it was not found", asset)
	}

	if err := DeleteAsset(accessToken, assetId); err != nil {
		t.Fatalf("Error deleting asset: %v", err)
	}

	time.Sleep(5)

	assets, err = GetAssets(accessToken)
	if err != nil {
		t.Fatalf("Error getting assets after deletion: %v", err)
	}

	for _, a := range assets {
		if a.Name == asset.Name && a.Description == asset.Description && a.Price == asset.Price {
			t.Fatalf("Expected asset %s to be deleted, but it was found", asset.Name)
		}
	}
}
