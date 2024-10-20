package register

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	funcional "testHive/testing/functional"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateRandomUser() (string, string) {
	rand.Seed(time.Now().UnixNano())
	username := fmt.Sprintf("user%d", rand.Intn(1000000))
	password := fmt.Sprintf("pass%d", rand.Intn(1000000))
	return username, password
}

func RegisterUser(username, password string) (User, error) {
	url := fmt.Sprintf("%s:/api/register", funcional.Api_url)

	user := User{Username: username, Password: password}
	jsonData, err := json.Marshal(user)
	if err != nil {
		return user, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return user, fmt.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	return user, nil
}
