package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Restaurant struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func GetRestaurantAddress(restaurantID int) (string, error) {
	// Create the URL to fetch the restaurant
	url := fmt.Sprintf("http://localhost:8080/restaurants/%d", restaurantID)

	// Send a GET request to the restaurant service
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch restaurant: %v", err)
	}
	defer response.Body.Close()

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch restaurant: received status code %d", response.StatusCode)
	}

	// Decode the response body
	var responseData map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return "", fmt.Errorf("failed to decode restaurant response: %v", err)
	}

	return responseData["data"].(map[string]interface{})["address"].(string), nil
}
