package controllers

import (
	"context"
	"fmt"
	proto "fulfillment-service/proto"
	"fulfillment-service/service"
)

type Server struct {
	proto.UnimplementedFulfillmentServiceServer
}

func (s *Server) AddDeliveryExecutive(_ context.Context, req *proto.AddDERequest) (*proto.AddDEResponse, error) {
	location := req.GetLocation()                        // Get the location of the Delivery Executive
	id, err := service.CreateDeliveryExecutive(location) // Create a new Delivery Executive
	if err != nil {
		return nil, fmt.Errorf("failed to create Delivery Executive: %v", err)
	}
	return &proto.AddDEResponse{Message: fmt.Sprintf("Delivery Executive added with ID %d", id)}, nil
}

func (s *Server) AssignDE(_ context.Context, req *proto.AssignDERequest) (*proto.AssignDEResponse, error) {
	orderID := req.GetOrderId()                         // Get the order ID
	deliveryExecutiveId := req.GetDeliveryExecutiveId() // Get the Delivery Executive ID

	err := service.AssignDeliveryExecutive(orderID, deliveryExecutiveId) // Assign the Delivery Executive to the order
	if err != nil {
		return nil, fmt.Errorf("failed to assign Delivery Executive: %v", err)
	}

	return &proto.AssignDEResponse{Message: fmt.Sprintf("Delivery Executive assigned to order %d", orderID)}, nil
}
