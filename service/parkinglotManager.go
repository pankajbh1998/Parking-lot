package service

import (
	"Parking-lot/models"
	"Parking-lot/repo"
	"fmt"
	"strings"
)

type PLManager struct {
	pl map[string]*Floor
	t *repo.Ticket
	v *repo.Vehicle
	size int64
}

func NewParkingLotManager(f []*Floor, plNames []string, t *repo.Ticket, v *repo.Vehicle)*PLManager{
	mp := make(map[string]*Floor)
	for i,val := range f {
		plNames[i] = strings.ToUpper(plNames[i])
		mp[plNames[i]]=val
	}
	return &PLManager{
		pl: mp,
		size: int64(len(plNames)),
		t: t,
		v: v,
	}
}

func (p *PLManager)AddFloorsToPL(plName string, slotsSize int64){
	p.pl[plName].AddFloor(slotsSize)
}

func(p *PLManager)getVehicle(vehicleType string, RegNo string, Color string)(veh *models.Vehicle){
	veh = p.v.GetVehicleByRegNo(RegNo)
	if veh != nil {
		return
	}
	veh = &models.Vehicle{
		Type: vehicleType,
		RegNo: RegNo,
		Color: Color,
	}
	p.v.AddVehicle(veh)
	return
}
func (p *PLManager)ParkVehicle(plName string, vehicleType string, RegNo string, Color string){
	vehicle := p.getVehicle(vehicleType, RegNo, Color)
	floorID, slotID, ok := p.pl[plName].Park(vehicle)
	if !ok {
		fmt.Println("Parking Lot Full")
		return
	}
	ticket := &models.Ticket{
		ParkingLotID: plName,
		FloorNo: floorID,
		SlotNo: slotID,
	}
	p.t.GetNewTicket(ticket)
	fmt.Printf("Parked vehicle. Ticket ID: %v\n",ticket.ConvertToStr())
	return
}


func (p *PLManager)UnParkVehicle(plName string, ticketID string){
	ticket := p.t.GetTicketByTicketFormat(ticketID)
	if ticket == nil {
		fmt.Println("Invalid Ticket")
		return
	}
	vehicleID, ok := p.pl[plName].UnPark(ticket.FloorNo, ticket.SlotNo)
	if !ok {
		fmt.Println("Invalid Ticket")
		return
	}
	vehicle := p.v.GetVehicle(vehicleID)
	fmt.Printf("Unparked vehicle with Registration Number: %v and Color: %v\n",vehicle.RegNo, vehicle.Color)
	p.t.RemoveTicket(ticket.ID)
	return
}

func (f *PLManager)GetOccupiedOfEmptySlots(plName string, slotType string){
	ret := f.pl[plName].GetOccupiedSlots(slotType)
	PrintSlots(ret, "Occupied", slotType)
	return
}

func (f *PLManager)GetEmptySlots(plName string, slotType string){
	ret := f.pl[plName].GetEmptySlots(slotType)
	PrintSlots(ret, "Free", slotType)
	return
}

func (f *PLManager)GetNumberOfFreeEmptySlots(plName string, slotType string){
	ret := f.pl[plName].GetNumberOfEmptySlots(slotType)
	//fmt.Println("Print Slots count")

	for i, slots := range ret {
		fmt.Printf("No. of free slots for %v on floor %v: %v\n", slotType, i+1, slots)
	}
	return
}



func PrintSlots(ret [][]int64, slotAvailability string, slotType string){
	//fmt.Println("Print Slots")
	for i, slots := range ret {
		str := fmt.Sprintf("%v slots for %v on floor %v [ ",slotAvailability, slotType,  i+1)
		for _,val := range slots {
			str += fmt.Sprintf("%v, ", val+1)
		}
		str += "]"
		fmt.Println(str)
	}
}

