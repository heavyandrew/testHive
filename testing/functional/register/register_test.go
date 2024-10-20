package register

import (
	"testing"
)

func TestRegisterUser(t *testing.T) {
	username, password := GenerateRandomUser()
	_, err := RegisterUser(username, password)
	if err != nil {
		t.Fatalf("Error registering user: %v", err)
	}
}
