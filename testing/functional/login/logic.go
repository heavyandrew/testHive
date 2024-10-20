package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	funcional "testHive/testing/functional"
	"testHive/testing/functional/register"
)

type UserResponse struct {
	AccessToken string `json:"access_token"`
}

func LoginUser(username, password string) (string, error) {
	url := fmt.Sprintf("%s:/api/login", funcional.Api_url)

	credentials := register.User{Username: username, Password: password}
	jsonData, err := json.Marshal(credentials)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var userResponse UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return "", err
	}

	if userResponse.AccessToken == "" {
		return "", fmt.Errorf("expected access token is empty")
	}

	return userResponse.AccessToken, nil
}
