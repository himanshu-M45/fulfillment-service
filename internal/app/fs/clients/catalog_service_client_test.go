package clients

import (
	"testing"
)

func TestGetRestaurantAddress_Success(t *testing.T) {
	result, err := GetRestaurantAddress(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != "123 Main St" {
		t.Fatalf("Expected '123 Main St', got %v", result)
	}
}

func TestGetRestaurantAddress_RequestError(t *testing.T) {
	_, err := GetRestaurantAddress(19)
	if err == nil || err.Error() != "failed to fetch restaurant: received status code 404" {
		t.Fatalf("Expected request error, got %v", err)
	}
}

func TestGetRestaurantAddress_StatusCodeError(t *testing.T) {
	_, err := GetRestaurantAddress(100)
	if err == nil || err.Error() != "failed to fetch restaurant: received status code 404" {
		t.Fatalf("Expected status code error, got %v", err)
	}
}
