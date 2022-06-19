package main

import (
	"Parking-lot/repo"
	"Parking-lot/service"
)

func initManagers(plName []string, floors int64, SlotNo int64)*service.PLManager{
	fManagers := make([]*service.Floor,0)
	for range plName{
		sRepos := make([]*repo.Slot,0)
		for i:=int64(0);i<floors;i++ {
			sRepos = append(sRepos, repo.NewSlot(SlotNo))
		}
		fManagers = append(fManagers, service.NewFloorManager(sRepos))
	}

	vRepo := repo.NewVehicle()
	tRepo := repo.NewTicket()
	return service.NewParkingLotManager(fManagers, plName, tRepo, vRepo)
}

func getCommands()[]string{
	commands := []string{
		"display free_count CAR",
		"display free_count BIKE",
		"display free_count TRUCK",
		"display free_slots CAR",
		"display free_slots BIKE",
		"display free_slots TRUCK",
		"display occupied_slots CAR",
		"display occupied_slots BIKE",
		"display occupied_slots TRUCK",
		"park_vehicle CAR KA-01-DB-1234 black",
		"park_vehicle CAR KA-02-CB-1334 red",
		"park_vehicle CAR KA-01-DB-1133 black",
		"park_vehicle CAR KA-05-HJ-8432 white",
		"park_vehicle CAR WB-45-HO-9032 white",
		"park_vehicle CAR KA-01-DF-8230 black",
		"park_vehicle CAR KA-21-HS-2347 red",
		"display free_count CAR",
		"display free_count BIKE",
		"display free_count TRUCK",
		"unpark_vehicle PR1234_2_5",
		"unpark_vehicle PR1234_2_5",
		"unpark_vehicle PR1234_2_7",
		"display free_count CAR",
		"display free_count BIKE",
		"display free_count TRUCK",
		"display free_slots CAR",
		"display free_slots BIKE",
		"display free_slots TRUCK",
		"display occupied_slots CAR",
		"display occupied_slots BIKE",
		"display occupied_slots TRUCK",
		"park_vehicle BIKE KA-01-DB-1541 black",
		"park_vehicle TRUCK KA-32-SJ-5389 orange",
		"park_vehicle TRUCK KL-54-DN-4582 green",
		"park_vehicle TRUCK KL-12-HF-4542 green",
		"display free_count CAR",
		"display free_count BIKE",
		"display free_count TRUCK",
		"display free_slots CAR",
		"display free_slots BIKE",
		"display free_slots TRUCK",
		"display occupied_slots CAR",
		"display occupied_slots BIKE",
		"display occupied_slots TRUCK",
	}
	return commands
}


func main(){
	plName := []string{"PR1234"}
	floorNo := int64(2)
	slotNo := int64(6)
	plManager := initManagers(plName, floorNo,slotNo)
	commands := getCommands()
	for _,command := range commands {
		RunCommand(command, plManager, plName[0])
	}
}

//Output
//No. of free slots for CAR on floor 1: 0
//No. of free slots for CAR on floor 2: 1
//No. of free slots for BIKE on floor 1: 2
//No. of free slots for BIKE on floor 2: 2
//No. of free slots for TRUCK on floor 1: 1
//No. of free slots for TRUCK on floor 2: 1
//Free slots for CAR on floor 1 [ ]
//Free slots for CAR on floor 2 [ 5, ]
//Free slots for BIKE on floor 1 [ 2, 3, ]
//Free slots for BIKE on floor 2 [ 2, 3, ]
//Free slots for TRUCK on floor 1 [ 1, ]
//Free slots for TRUCK on floor 2 [ 1, ]
//Occupied slots for CAR on floor 1 [ 4, 5, 6, ]
//Occupied slots for CAR on floor 2 [ 4, 6, ]
//Occupied slots for BIKE on floor 1 [ ]
//Occupied slots for BIKE on floor 2 [ ]
//Occupied slots for TRUCK on floor 1 [ ]
//Occupied slots for TRUCK on floor 2 [ ]
//Parked vehicle. Ticket ID: PR1234_1_2
//Parked vehicle. Ticket ID: PR1234_1_1
//Parked vehicle. Ticket ID: PR1234_2_1
//Parking Lot Full
//No. of free slots for CAR on floor 1: 0
//No. of free slots for CAR on floor 2: 1
//No. of free slots for BIKE on floor 1: 1
//No. of free slots for BIKE on floor 2: 2
//No. of free slots for TRUCK on floor 1: 0
//No. of free slots for TRUCK on floor 2: 0
//Free slots for CAR on floor 1 [ ]
//Free slots for CAR on floor 2 [ 5, ]
//Free slots for BIKE on floor 1 [ 3, ]
//Free slots for BIKE on floor 2 [ 2, 3, ]
//Free slots for TRUCK on floor 1 [ ]
//Free slots for TRUCK on floor 2 [ ]
//Occupied slots for CAR on floor 1 [ 5, 6, 4, ]
//Occupied slots for CAR on floor 2 [ 4, 6, ]
//Occupied slots for BIKE on floor 1 [ 2, ]
//Occupied slots for BIKE on floor 2 [ ]
//Occupied slots for TRUCK on floor 1 [ 1, ]
//Occupied slots for TRUCK on floor 2 [ 1, ]
//Exiting.


