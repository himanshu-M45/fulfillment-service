package main

import (
	"context"
	"fmt"
	"fulfillment-service/db"
	proto "fulfillment-service/proto"
)

type Server struct {
	proto.UnimplementedFulfillmentServiceServer
}

func (s *Server) AddDeliveryExecutive(_ context.Context, req *proto.AddDERequest) (*proto.AddDEResponse, error) {
	location := req.GetLocation() // Get the location of the Delivery Executive
	database := db.GetDB()        // Get the database connection
	if database == nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}
	// Insert new Delivery Executive into the database
	query := "INSERT INTO delivery_executives (location, is_available, assigned_order_id) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := database.QueryRow(query, location, true, -1).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to add Delivery Executive: %v", err)
	}

	return &proto.AddDEResponse{Message: fmt.Sprintf("Delivery Executive added with ID %d", id)}, nil
}
