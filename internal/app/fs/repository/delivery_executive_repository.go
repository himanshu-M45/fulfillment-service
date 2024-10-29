package repository

import (
	"fmt"
	"fulfillment-service/internal/app/fs/db"
	"fulfillment-service/internal/app/fs/models"
)

// SaveDeliveryExecutive inserts a new Delivery Executive into the database
// 1. Get the database connection
// 2. Create a new Delivery Executive with the provided location
// 3. Insert the Delivery Executive into the database
// 4. Return the ID of the newly created Delivery Executive
func SaveDeliveryExecutive(location string) (int, error) {
	database := db.GetDB()
	if database == nil {
		return -1, fmt.Errorf("failed to get database connection")
	}
	deliveryExecutive := models.DeliveryExecutive{
		Location:        location,
		IsAvailable:     true,
		AssignedOrderId: -1,
	}
	result := database.Create(&deliveryExecutive)
	if result.Error != nil {
		return 0, result.Error
	}
	return deliveryExecutive.ID, nil
}

// GetAllDeliveryExecutives retrieves all Delivery Executives from the database
// 1. Get the database connection
// 2. Query the database for all Delivery Executives
// 3. Return the list of Delivery Executives
func GetAllDeliveryExecutives() ([]models.DeliveryExecutive, error) {
	database := db.GetDB()
	if database == nil {
		return nil, fmt.Errorf("failed to get database connection")
	}
	var deliveryExecutives []models.DeliveryExecutive
	result := database.Find(&deliveryExecutives)
	if result.Error != nil {
		return nil, result.Error
	}
	return deliveryExecutives, nil
}

func GetDeliveryExecutive(deliveryExecutiveId int) (bool, int, error) {
	database := db.GetDB()
	if database == nil {
		return false, 0, fmt.Errorf("failed to get database connection")
	}
	var deliveryExecutive models.DeliveryExecutive
	result := database.First(&deliveryExecutive, deliveryExecutiveId)
	if result.Error != nil {
		return false, 0, result.Error
	}
	return deliveryExecutive.IsAvailable, deliveryExecutive.AssignedOrderId, nil
}

// UpdateDeliveryExecutiveStatus updates the status of a Delivery Executive
// 1. Get the database connection
// 2. Update the Delivery Executive's availability and assigned order ID
// 3. Return any errors that occur during the update
func UpdateDeliveryExecutiveStatus(isAvailable bool, orderID int, deliveryExecutiveId int) error {
	database := db.GetDB()
	if database == nil {
		return fmt.Errorf("failed to get database connection")
	}
	result := database.Model(&models.DeliveryExecutive{}).Where("id = ?", deliveryExecutiveId).Updates(map[string]interface{}{
		"is_available":      isAvailable,
		"assigned_order_id": orderID,
	})
	return result.Error
}
