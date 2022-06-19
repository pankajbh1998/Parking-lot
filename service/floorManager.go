package service

import (
	"Parking-lot/models"
	"Parking-lot/repo"
)

type Floor struct {
	mp map[int64]*repo.Slot
	size int64
}

func NewFloorManager(s []*repo.Slot)*Floor{
	mp := make(map[int64]*repo.Slot)
	for i,val := range s {
		mp[int64(i)]=val
	}
	return &Floor{
		mp: mp,
		size: int64(len(s)),
	}
}

func (f *Floor)AddFloor(slotSize int64){
	f.mp[f.size] = repo.NewSlot(slotSize)
	f.size++
}

func (f *Floor)GetEmptySlots(slotType string)[][]int64{
	ret := make([][]int64, f.size)
	for i:=int64(0);i<f.size;i++ {
		ret[i] = f.mp[i].GetAllEmptyByType(slotType)
	}
	return ret
}

func (f *Floor)GetOccupiedSlots(slotType string)[][]int64{
	ret := make([][]int64, f.size)
	for i:=int64(0);i<f.size;i++ {
		ret[i]=f.mp[i].GetAllOccupiedSlotByType(slotType)
	}
	return ret
}

func (f *Floor)GetNumberOfEmptySlots(slotType string)[]int{
	ret := make([]int, f.size)
	for i:=int64(0);i<f.size;i++ {
		ret[i] = len(f.mp[i].GetAllEmptyByType(slotType))
	}
	return ret
}

func (f *Floor)UnPark(floorID int64, slotID int64)(int64, bool){
	return f.mp[floorID].ReleaseSlot(slotID)
}

func (f *Floor)Park(vehicle *models.Vehicle)(int64, int64, bool){
	for i:=int64(0);i<f.size;i++ {
		id := f.mp[i].FindNearestAvailableSlotID(vehicle.Type)
		if id != -1{
			f.mp[i].BookSlot(id, vehicle.ID)
			return i, id, true
		}
	}
	return -1, -1, false
}

