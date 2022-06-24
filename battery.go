package main

import "math"

var columnID int = 1
var floorRequestButtonID int = 1
var floor int = 1

type Battery struct {
	ID                      int
	status                  string
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
	servedFloors            []int
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	b := new(Battery)
	b.ID = _id
	b.status = "online"

	if _amountOfBasements > 0 {
		b.createBasementFloorRequestButtons(_amountOfBasements)
		b.createBasmentColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	b.createFloorRequestButtons(_amountOfFloors)
	b.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)

	return b
}

func (b *Battery) createBasmentColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	var servedFloors []int
	floor = -1

	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	column := NewColumn(columnID, "online", _amountOfBasements, _amountOfElevatorPerColumn, servedFloors, true)
	b.columnsList = append(b.columnsList, *column)
	columnID--
}

func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevatorPerColumn int) {
	var amountOfFloorsPerColumn = int(math.Round(float64(_amountOfFloors / _amountOfColumns)))
	floor = 1

	for i := 0; i < _amountOfColumns; i++ {
		var servedFloors []int
		for y := 0; y < amountOfFloorsPerColumn; y++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := NewColumn(columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		columnID++
	}
}

func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	var buttonFloor int = 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, "OFF", buttonFloor, "Up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	var buttonFloor int = -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, "OFF", buttonFloor, "Down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor) {
			return &column
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := *b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return &column, elevator
}
