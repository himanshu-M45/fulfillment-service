package clients

import (
	"encoding/json"
	"fmt"
	proto "fulfillment-service/proto"
	"net/http"
)

// Order represents the structure of the order data
type Order struct {
	ID              int         `json:"id"`
	UserID          int         `json:"userId"`
	RestaurantID    int         `json:"restaurantId"`
	DeliveryAddress string      `json:"deliveryAddress"`
	TotalPrice      int         `json:"totalPrice"`
	Status          OrderStatus `json:"status"`
}

// OrderStatus represents the status of an order
type OrderStatus string

const (
	DE_ALLOCATED     OrderStatus = "DE_ALLOCATED"
	OUT_FOR_DELIVERY OrderStatus = "OUT_FOR_DELIVERY"
	DELIVERED        OrderStatus = "DELIVERED"
)

// CheckOrderCredibility checks if an order is available for DE assignment
// 1. Create the URL to fetch the order
// 2. Send a GET request to the order service
// 3. Check if the request was successful
// 4. Decode the response body
// 5. Check if the order is available for DE assignment
// 6. Return the response message
func CheckOrderCredibility(orderId int) (string, error) {
	url := fmt.Sprintf("http://localhost:8081/orders/%d", orderId)
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch order: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch order: received status code %d", response.StatusCode)
	}

	var responseData map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return "", fmt.Errorf("failed to decode order response: %v", err)
	}

	data := responseData["data"].(map[string]interface{})
	order := Order{
		ID:              int(data["orderId"].(float64)),
		UserID:          int(data["userId"].(float64)),
		RestaurantID:    int(data["restaurantId"].(float64)),
		DeliveryAddress: data["deliveryAddress"].(string),
		TotalPrice:      int(data["price"].(float64)),
		Status:          OrderStatus(data["status"].(string)),
	}

	if order.Status != "ORDER_CREATED" {
		return "", fmt.Errorf("order is not available for DE assignment")
	}

	return "order can be assigned", nil
}

// UpdateOrderStatus updates the status of an order
// 1. Create the URL to update the order status
// 2. Create a PUT request to update the order status
// 3. Send the request to the order service
// 4. Check if the request was successful
// 5. Return the response message
func UpdateOrderStatus(orderId int, status OrderStatus) (*proto.UpdateStatusResponse, error) {
	url := fmt.Sprintf("http://localhost:8081/orders/%d?status=%s", orderId, status)

	request, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update order status, received status code %d", response.StatusCode)
	}

	return &proto.UpdateStatusResponse{Message: "Order status updated successfully"}, nil
}
