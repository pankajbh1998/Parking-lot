package repo

import "Parking-lot/models"

type Slot struct {
	slots map[int64]*models.Slot
	totalSlots int64
}

func NewSlot(size int64)*Slot {
	slots := make(map[int64]*models.Slot,0)
	slots[0] = &models.Slot{
		ID: 0,
		Type: models.SlotTypeTruck,
		IsFilled: false,
	}

	for i:=int64(1);i<3;i++ {
		slots[i]=&models.Slot{
			ID: i,
			Type: models.SlotTypeBike,
		}
	}

	for i:=int64(3);i<size;i++ {
		slots[i]=&models.Slot{
			ID: i,
			Type: models.SlotTypeCar,
		}
	}
	return &Slot{
		slots: slots,
		totalSlots: size,
	}
}

func (s *Slot)BookSlot(id int64, vehicleID int64){
	s.slots[id].IsFilled = true
	s.slots[id].VehicleID = vehicleID
}

func (s *Slot)GetAllOccupiedSlotByType(slotType string)[]int64{
	ret := make([]int64,0)
	for _, val := range s.slots {
		if val.Type == slotType && val.IsFilled {
			ret = append(ret, val.ID)
		}
	}
	return ret
}

func (s *Slot)GetAllEmptyByType(slotType string)[]int64{
	ret := make([]int64,0)
	for _, val := range s.slots {
		if val.Type == slotType && !val.IsFilled {
			ret = append(ret, val.ID)
		}
	}
	return ret
}

func (s *Slot)FindNearestAvailableSlotID(vType string)int64{
	for i:=int64(0);i<s.totalSlots;i++ {
		val ,_ := s.slots[i]
		if val.Type == vType && !val.IsFilled {
			return i
		}
	}
	return -1
}

func (s *Slot)ReleaseSlot(id int64)(int64,bool){
	slot, ok := s.slots[id]
	if !ok || !slot.IsFilled {
		return -1, false
	}
	slot.IsFilled = false
	return slot.VehicleID, true
}




