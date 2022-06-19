package repo

import (
	"Parking-lot/models"
	"math/rand"
)

type Vehicle struct {
	mp map[int64]*models.Vehicle
	mpRegNo map[string]*models.Vehicle
}

func NewVehicle()*Vehicle{
	return &Vehicle{
		mp: make(map[int64]*models.Vehicle),
		mpRegNo: make(map[string]*models.Vehicle),
	}
}

func (v *Vehicle)AddVehicle(vehicle *models.Vehicle) int64 {
	vehicle.ID = rand.Int63()
	v.mp[vehicle.ID] = vehicle
	v.mpRegNo[vehicle.RegNo] = vehicle
	return vehicle.ID
}

func (v *Vehicle)GetVehicle(id int64)*models.Vehicle {
	val, ok := v.mp[id]
	if !ok {
		return nil
	}
	return val
}

func (v *Vehicle)GetVehicleByRegNo(id string)*models.Vehicle {
	val, ok := v.mpRegNo[id]
	if !ok {
		return nil
	}
	return val
}



