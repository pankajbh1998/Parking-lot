package models

import "fmt"

type Ticket struct {
	ID  int64
	ParkingLotID string
	FloorNo int64
	SlotNo int64
}

func (ticket *Ticket) ConvertToStr()string{
	return fmt.Sprintf("%v_%v_%v",ticket.ParkingLotID, ticket.FloorNo+1, ticket.SlotNo+1);
}