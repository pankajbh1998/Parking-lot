package main

import (
	"Parking-lot/models"
	"Parking-lot/service"
	"strings"
)

func RunCommand(command string, manager *service.PLManager, plName string){
	str := strings.Split(command, " ")
	switch strings.ToLower(str[0]){
		case models.DisplayCommand:
			disPlayCommand(str[1], str[2], manager, plName)
		case models.ParkVehicle:
			parkVehicle(str[1], str[2], str[3], manager, plName)
		case models.UnParkVehicle:
			unParkVehicle(str[1], manager, plName)
	}
}

func unParkVehicle(ticket string, manager *service.PLManager, name string) {
	manager.UnParkVehicle(name, ticket)
}

func parkVehicle(vehicleType string, RegNo string, Color string, manager *service.PLManager, name string) {
	manager.ParkVehicle(name, vehicleType, RegNo, Color)
}

func disPlayCommand(displayType string, slotType string, manager *service.PLManager, name string) {
	switch displayType {
	case models.FreeCount:
		manager.GetNumberOfFreeEmptySlots(name,slotType)
	case models.FreeSlots:
		manager.GetEmptySlots(name, slotType)
	case models.OccupiedSlots:
		manager.GetOccupiedOfEmptySlots(name, slotType)
	}
}



