package service

import (
	"fmt"
	"fulfillment-service/internal/app/fs/clients"
	"fulfillment-service/internal/app/fs/repository"
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
// 1. Check the credibility of the order
// 2. Get the Delivery Executive details
// 3. Check if the Delivery Executive is available and not assigned to any order
// 4. Update the Delivery Executive status
// 5. Update the Order status
func AssignDeliveryExecutive(orderID int, deliveryExecutiveId int) error {
	// Check Order credibility
	message, err := clients.CheckOrderCredibility(orderID)
	if err != nil && message != "order can be assigned" {
		return fmt.Errorf("failed to check order credibility: %v", err)
	}

	// Get Delivery Executive details
	isAvailable, assignedOrderID, err := repository.GetDeliveryExecutive(deliveryExecutiveId)

	// Check if Delivery Executive is available and not assigned to any order
	if !isAvailable || assignedOrderID != -1 {
		return fmt.Errorf("delivery Executive is not available or already assigned to an order")
	}

	// Update Delivery Executive status
	err = repository.UpdateDeliveryExecutiveStatus(false, orderID, deliveryExecutiveId)
	if err != nil {
		return fmt.Errorf("failed to update Delivery Executive status: %v", err)
	}

	// Update Order status
	_, err = clients.UpdateOrderStatus(orderID, clients.DE_ALLOCATED)
	if err != nil {
		return fmt.Errorf("failed to update order status: %v", err)
	}

	return nil
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
