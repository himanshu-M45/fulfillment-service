package clients

import (
	"testing"
)

func TestCheckOrderCredibility_Success(t *testing.T) {
	result, err := CheckOrderCredibility(21)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != "order can be assigned" {
		t.Fatalf("Expected 'order can be assigned', got %v", result)
	}
}

func TestCheckOrderCredibility_FetchError(t *testing.T) {
	_, err := CheckOrderCredibility(18)
	if err == nil || err.Error() != "failed to fetch order: Get \"http://localhost:8081/orders/18\": dial tcp [::1]:8081: connect: connection refused" {
		t.Fatalf("Expected fetch error, got %v", err)
	}
}

func TestCheckOrderCredibility_StatusCodeError(t *testing.T) {
	_, err := CheckOrderCredibility(180)
	if err == nil || err.Error() != "failed to fetch order: received status code 404" {
		t.Fatalf("Expected status code error, got %v", err)
	}
}

func TestCheckOrderCredibility_DecodeError(t *testing.T) {
	_, err := CheckOrderCredibility(21)
	if err == nil || err.Error() != "failed to decode order response: invalid character 'i' looking for beginning of value" {
		t.Fatalf("Expected decode error, got %v", err)
	}
}

func TestCheckOrderCredibility_OrderNotAvailable(t *testing.T) {
	_, err := CheckOrderCredibility(18)
	if err == nil || err.Error() != "order is not available for DE assignment" {
		t.Fatalf("Expected order not available error, got %v", err)
	}
}

// ----------------------------------------------------------
func TestUpdateOrderStatus_Success(t *testing.T) {
	response, err := UpdateOrderStatus(18, OutForDelivery)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response.Message != "Order status updated successfully" {
		t.Fatalf("Expected 'Order status updated successfully', got %v", response.Message)
	}
}

func TestUpdateOrderStatus_RequestError(t *testing.T) {
	_, err := UpdateOrderStatus(18, DELIVERED)
	if err == nil ||
		err.Error() != "failed to update order status: Put \"http://localhost:8081/orders/18?status=DELIVERED\": dial tcp [::1]:8081: connect: connection refused" {
		t.Fatalf("Expected request error, got %v", err)
	}
}

func TestUpdateOrderStatus_UpdateError(t *testing.T) {
	_, err := UpdateOrderStatus(18, DELIVERED)
	if err == nil || err.Error() != "failed to update order status: received status code 400" {
		t.Fatalf("Expected update error, got %v", err)
	}
}
