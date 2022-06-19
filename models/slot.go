package models

type Slot struct {
	ID       int64
	Type     string
	IsFilled bool
	VehicleID int64
}
