package model

type DeliveryExecutive struct {
	ID              int
	Location        string
	IsAvailable     bool
	AssignedOrderId int
}
