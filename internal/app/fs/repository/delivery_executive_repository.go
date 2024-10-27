package repository

import (
	"database/sql"
	"fmt"
	"fulfillment-service/internal/app/fs/db"
)

// getDBConnection retrieves the database connection and handles the error
// 1. Get the database connection
// 2. Return the database connection or an error
func getDBConnection() (*sql.DB, error) {
	database := db.GetDB()
	if database == nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}
	return database, nil
}

// SaveDeliveryExecutive inserts a new Delivery Executive into the database
// 1. Get the database connection
// 2. Prepare the query to insert the Delivery Executive
// 3. Execute the query and return the ID of the newly inserted Delivery Executive
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
// 1. Get the database connection
// 2. Prepare the query to retrieve the Delivery Executive
// 3. Execute the query and return the Delivery Executive
func GetDeliveryExecutive(deliveryExecutiveId int) (bool, int, error) {
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
// 1. Get the database connection
// 2. Prepare the query to update the Delivery Executive status
// 3. Execute the query and return the result
func UpdateDeliveryExecutiveStatus(isAvailable bool, orderID int, deliveryExecutiveId int) error {
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
