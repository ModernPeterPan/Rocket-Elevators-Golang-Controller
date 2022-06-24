package main

import "math"

var elevatorID int = 1
var callButtonID int = 1

type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	elevatorsList     []*Elevator
	callButtonsList   []*CallButton
	servedFloorsList  []int
	elevator          Elevator
}

func NewColumn(_id int, _status string, _amountOfFloors int, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	c := new(Column)
	c.ID = _id
	c.status = _status
	c.servedFloorsList = _servedFloors

	c.createElevators(_amountOfFloors, _amountOfElevators)
	c.createCallButtons(_amountOfFloors, _isBasement)
	return c
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		var buttonFloor int = -1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := NewCallButton(callButtonID, "OFF", buttonFloor, "Up")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor--
			callButtonID++
		}
	} else {
		var buttonFloor int = 1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := NewCallButton(callButtonID, "OFF", buttonFloor, "Down")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor++
			callButtonID++
		}
	}
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	for i := 0; i < _amountOfElevators; i++ {
		c.elevatorsList = append(c.elevatorsList, NewElevator(elevatorID, "idle", _amountOfFloors, 1, ""))
		elevatorID++
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(userPosition int, _direction string) *Elevator {
	elevator := c.findElevator(userPosition, _direction)
	elevator.addNewRequest(userPosition)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

func (c *Column) findElevator(userPosition int, _requestedDirection string) *Elevator {
	var bestElevator *Elevator
	var bestScore int = 6
	var referenceGap int = 100000

	if userPosition == 1 {
		for _, elevator := range c.elevatorsList {
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, userPosition)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if userPosition == elevator.currentFloor && elevator.status == "stopped" && _requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if userPosition > elevator.currentFloor && elevator.direction == "up" && _requestedDirection == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if userPosition < elevator.currentFloor && elevator.direction == "down" && _requestedDirection == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, userPosition)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, userPosition)
			}
		}
	}
	return bestElevator
}

func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, userPosition int) (*Elevator, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor) - float64(userPosition)))
	} else if bestScore == scoreToCheck {
		var gap int = int(math.Abs(float64(newElevator.currentFloor) - float64(userPosition)))
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	return bestElevator, bestScore, referenceGap
}
