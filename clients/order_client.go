package clients

import (
	"encoding/json"
	"fmt"
	proto "fulfillment-service/proto"
	"net/http"
)

type OrderClient interface {
	CheckOrderCredibility(orderID int32) (string, error)
	UpdateOrderStatus(orderID int32, status string) (string, error)
}

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
	DeAllocated    OrderStatus = "DE_ALLOCATED"
	OutForDelivery OrderStatus = "OUT_FOR_DELIVERY"
	DELIVERED      OrderStatus = "DELIVERED"
)

// CheckOrderCredibility checks if an order is available for DE assignment
func CheckOrderCredibility(orderId int32) (string, error) {
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

func UpdateOrderStatus(orderId int32, status OrderStatus) (*proto.UpdateStatusResponse, error) {
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
		return nil, fmt.Errorf("failed to update order status: received status code %d", response.StatusCode)
	}

	return &proto.UpdateStatusResponse{Message: "Order status updated successfully"}, nil
}
