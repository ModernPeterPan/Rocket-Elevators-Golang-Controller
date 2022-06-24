package main

import "sort"

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
	screenDisplay         int
	direction             string
	overweight            bool
}

func NewElevator(_id int, _status string, _amountOfFloors int, _currentFloor int, _elevatorID string) *Elevator {
	e := new(Elevator)
	e.ID = _id
	e.status = _status
	e.amountOfFloors = _amountOfFloors
	e.currentFloor = _currentFloor
	e.direction = ""
	e.door = Door{1, ""}
	return e
}

func (e *Elevator) move() {
	for len(e.floorRequestsList) > 0 {
		var destination int = e.floorRequestsList[0]
		e.status = "moving"
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		e.floorRequestsList = e.floorRequestsList[1:]
		e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
	}
	e.status = "idle"
	e.direction = ""
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)
	} else if e.direction == "down" {
		sort.Ints(e.floorRequestsList)
		e.floorRequestsList = (e.floorRequestsList)
	}
}

func (e *Elevator) operateDoors() {
	var obstruction bool = false
	e.door.status = "opened"
	if !e.overweight {
		e.door.status = "closing"
		if !obstruction {
			e.door.status = "closed"
		} else {
			e.operateDoors()
		}
	} else {
		for e.overweight {
			//Activate overweight alarm
		}
		e.operateDoors()
	}
}

func (e *Elevator) addNewRequest(resquestedFloor int) {
	if !contains(e.floorRequestsList, resquestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, resquestedFloor)
	}
	if e.currentFloor < resquestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > resquestedFloor {
		e.direction = "down"
	}
}
