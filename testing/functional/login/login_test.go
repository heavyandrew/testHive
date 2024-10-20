package login

import (
	"testHive/testing/functional/register"
	"testing"
	"time"
)

func TestRegisterAndLogin(t *testing.T) {
	username, password := register.GenerateRandomUser()
	_, err := register.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("Error registering user: %v", err)
	}
	time.Sleep(2)

	_, err = LoginUser(username, password)
	if err != nil {
		t.Fatalf("Error logging in: %v", err)
	}
}
