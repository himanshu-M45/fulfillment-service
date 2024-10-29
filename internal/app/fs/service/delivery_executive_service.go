package service

import (
	"fmt"
	"fulfillment-service/internal/app/fs/clients"
	"fulfillment-service/internal/app/fs/models"
	"fulfillment-service/internal/app/fs/repository"
	"strings"
)

// CreateDeliveryExecutive creates a new Delivery Executive into the database
// 1. Save the Delivery Executive details
// 2. Return the ID of the newly inserted Delivery Executive
func CreateDeliveryExecutive(location string) (int, error) {
	id, err := repository.SaveDeliveryExecutive(location)
	if err != nil {
		return 0, fmt.Errorf("failed to add Delivery Executive: %v", err)
	}
	return id, nil
}

// AssignDeliveryExecutive assigns a Delivery Executive to an order
// 1. Check Order credibility
// 2. Get restaurant address
// 3. Get all Delivery Executives
// 4. Find a matching Delivery Executive
// 5. Update Delivery Executive status
// 6. Update Order status
// 7. Return the ID of the assigned Delivery Executive
func AssignDeliveryExecutive(orderID int) (int, error) {
	// Check Order credibility
	restaurantId, err := clients.CheckOrderCredibility(orderID)
	if err != nil {
		return -1, fmt.Errorf("failed to check order credibility: %v", err)
	}

	// Get restaurant address
	restaurantAddress, err := clients.GetRestaurantAddress(restaurantId)
	if err != nil {
		return -1, fmt.Errorf("failed to get restaurant address: %v", err)
	}

	// Get all Delivery Executives
	deliveryExecutives, err := repository.GetAllDeliveryExecutives()
	if err != nil {
		return -1, fmt.Errorf("failed to get delivery executives: %v", err)
	}

	// Find a matching Delivery Executive
	var selectedDE *models.DeliveryExecutive
	for _, de := range deliveryExecutives {
		if de.IsAvailable && de.AssignedOrderId == -1 && strings.Contains(restaurantAddress, de.Location) {
			selectedDE = &de
			break
		}
	}

	if selectedDE == nil {
		return -1, fmt.Errorf("no delivery executive available for the order")
	}

	// Update Delivery Executive status
	err = repository.UpdateDeliveryExecutiveStatus(false, orderID, selectedDE.ID)
	if err != nil {
		return -1, fmt.Errorf("failed to update Delivery Executive status: %v", err)
	}

	// Update Order status
	_, err = clients.UpdateOrderStatus(orderID, clients.DE_ALLOCATED)
	if err != nil {
		return -1, fmt.Errorf("failed to update order status: %v", err)
	}

	return selectedDE.ID, nil
}

// UpdateOrderStatus updates the status of an order
// 1. Check if the status is not DE_ALLOCATED
// 2. Extract order ID by delivery executive ID
// 3. Update order status
// 4. Update delivery executive status if the status is DELIVERED
func UpdateOrderStatus(deliveryExecutiveId int, status string) error {
	if clients.OrderStatus(status) == clients.DE_ALLOCATED {
		return fmt.Errorf("cannot update order status to DE_ALLOCATED")
	}

	// extract order id by delivery executive id
	_, orderID, err := repository.GetDeliveryExecutive(deliveryExecutiveId)
	if err != nil {
		return fmt.Errorf("failed to get order id: %v", err)
	}

	// update order status
	_, err = clients.UpdateOrderStatus(orderID, clients.OrderStatus(status))
	if err != nil {
		return fmt.Errorf("failed to update order status: %v", err)
	}

	if clients.OrderStatus(status) == clients.DELIVERED {
		// update delivery executive status
		err = repository.UpdateDeliveryExecutiveStatus(true, -1, deliveryExecutiveId)
		if err != nil {
			return fmt.Errorf("failed to update delivery executive status: %v", err)
		}
	}

	return nil
}
