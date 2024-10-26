package repository

import (
	"database/sql"
	"fmt"
	"fulfillment-service/db"
)

// getDBConnection retrieves the database connection and handles the error
func getDBConnection() (*sql.DB, error) {
	database := db.GetDB()
	if database == nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}
	return database, nil
}

// SaveDeliveryExecutive inserts a new Delivery Executive into the database
func SaveDeliveryExecutive(location string) (int, error) {
	database, err := getDBConnection()
	if err != nil {
		return 0, fmt.Errorf("failed to get database connection: %v", err)
	}

	query := "INSERT INTO delivery_executives (location, is_available, assigned_order_id) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err = database.QueryRow(query, location, true, -1).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to add Delivery Executive: %v", err)
	}
	return id, nil
}

// GetDeliveryExecutive retrieves a Delivery Executive by ID
func GetDeliveryExecutive(deliveryExecutiveId int32) (bool, int, error) {
	database, err := getDBConnection()
	if err != nil {
		return false, 0, fmt.Errorf("failed to get database connection: %v", err)
	}

	var isAvailable bool
	var assignedOrderID int
	query := "SELECT is_available, assigned_order_id FROM delivery_executives WHERE id = $1"
	err = database.QueryRow(query, deliveryExecutiveId).Scan(&isAvailable, &assignedOrderID)
	if err != nil {
		return false, 0, fmt.Errorf("failed to retrieve Delivery Executive: %v", err)
	}
	return isAvailable, assignedOrderID, nil
}

// UpdateDeliveryExecutiveStatus updates the status of a Delivery Executive
func UpdateDeliveryExecutiveStatus(isAvailable bool, orderID int32, deliveryExecutiveId int32) error {
	database, err := getDBConnection()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %v", err)
	}

	updateQuery := "UPDATE delivery_executives SET is_available = $1, assigned_order_id = $2 WHERE id = $3"
	_, err = database.Exec(updateQuery, isAvailable, orderID, deliveryExecutiveId)
	if err != nil {
		return fmt.Errorf("failed to update Delivery Executive status: %v", err)
	}
	return nil
}
