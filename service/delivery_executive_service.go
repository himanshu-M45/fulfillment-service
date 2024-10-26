package service

import (
	"fmt"
	"fulfillment-service/clients"
	"fulfillment-service/repository"
)

// CreateDeliveryExecutive creates a new Delivery Executive into the database
func CreateDeliveryExecutive(location string) (int, error) {
	id, err := repository.SaveDeliveryExecutive(location)
	if err != nil {
		return 0, fmt.Errorf("failed to add Delivery Executive: %v", err)
	}
	return id, nil
}

// AssignDeliveryExecutive assigns a Delivery Executive to an order
func AssignDeliveryExecutive(orderID int32, deliveryExecutiveId int32) error {
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
	_, err = clients.UpdateOrderStatus(orderID, clients.DeAllocated)
	if err != nil {
		return fmt.Errorf("failed to update order status: %v", err)
	}

	return nil
}
