package controllers

import (
	"context"
	"fmt"
	"fulfillment-service/internal/app/fs/service"
	proto "fulfillment-service/proto"
)

type Server struct {
	proto.UnimplementedFulfillmentServiceServer
}

// AddDeliveryExecutive adds a new Delivery Executive to the system
// 1. Get the location of the Delivery Executive
// 2. Create a new Delivery Executive
// 3. Return the ID of the newly created Delivery Executive
func (s *Server) AddDeliveryExecutive(_ context.Context, req *proto.AddDERequest) (*proto.AddDEResponse, error) {
	location := req.GetLocation()                        // Get the location of the Delivery Executive
	id, err := service.CreateDeliveryExecutive(location) // Create a new Delivery Executive
	if err != nil {
		return nil, fmt.Errorf("failed to create Delivery Executive: %v", err)
	}
	return &proto.AddDEResponse{Message: fmt.Sprintf("Delivery Executive added with ID %d", id)}, nil
}

// AssignDE assigns a Delivery Executive to an order
// 1. Get the order ID and Delivery Executive ID
// 2. Assign the Delivery Executive to the order
// 3. Return a success message
func (s *Server) AssignDE(_ context.Context, req *proto.AssignDERequest) (*proto.AssignDEResponse, error) {
	orderID := req.GetOrderId()                         // Get the order ID
	deliveryExecutiveId := req.GetDeliveryExecutiveId() // Get the Delivery Executive ID

	err := service.AssignDeliveryExecutive(int(orderID), int(deliveryExecutiveId)) // Assign the Delivery Executive to the order
	if err != nil {
		return nil, fmt.Errorf("failed to assign Delivery Executive: %v", err)
	}

	return &proto.AssignDEResponse{Message: fmt.Sprintf("Delivery Executive assigned to order %d", orderID)}, nil
}

// UpdateOrderStatus updates the status of an order
// 1. Get the Delivery Executive ID and the new status
// 2. Update the status of the order
// 3. Return a success message
func (s *Server) UpdateOrderStatus(_ context.Context, req *proto.UpdateStatusRequest) (*proto.UpdateStatusResponse, error) {
	deliveryExecutiveId := req.GetDeliveryExecutiveId() // Get the order ID
	status := req.GetStatus()                           // Get the status

	err := service.UpdateOrderStatus(int(deliveryExecutiveId), status) // Update the order status
	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %v", err)
	}

	return &proto.UpdateStatusResponse{Message: fmt.Sprintf("Order status updated to: %v", status)}, nil
}
