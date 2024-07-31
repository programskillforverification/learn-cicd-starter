package auth_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	url := "https://www.google.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", "My-App/1.0")
	req.Header.Set("Authorization", "ApiKey duhfuf889")

	result, err := auth.GetAPIKey(req.Header)
	expected := "duhfuf889"
	if err != nil {
		fmt.Println("Get API fail:", err)
		return
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
