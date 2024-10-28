package models

type DeliveryExecutive struct {
	ID              int `gorm:"primaryKey;autoIncrement"`
	Location        string
	IsAvailable     bool
	AssignedOrderId int
}
