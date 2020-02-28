package main

import (
	"fmt"
	"sort"
	"time"
)

//RequestElevator and AssignElevator

//FloorNumber = the place where the customer is at
//RequestedFloor = the floor the customer want to go

// ElevatorController hold
type ElevatorController struct {
	BatteryNumber 	int
	batteries       []Battery
	nbColumn  	 	int
	userMove  		string
}

// Battery hold
type Battery struct {
	nbColumn int
	columnList   []Column
}

// Column hold
type Column struct {
	ColumnNumber    int
	nbEleInColumn 	int
	ElevatorList    []Elevator
}

// Elevator hold
type Elevator struct {
	eleNumber          int
	currentFloor	   int
	floorList          []int
	ElevatorStatus     string
	ElevatorDirection  string
	doorSensor         bool
	Column             Column
}

// NewController is where Battery are create

func NewController(BatteryNumber int) ElevatorController {
	controller := new(ElevatorController)
	controller.BatteryNumber = 1
	for index := 0; index < BatteryNumber; index++ {
		battery := NewBattery(index)
		controller.batteries = append(controller.batteries, *battery)
	}
	return *controller
}

// NewBattery is where Column are create

func NewBattery(nbColumn int) *Battery {
	battery := new(Battery)
	battery.nbColumn = 4
	for index := 0; index < battery.nbColumn; index++ {
		Column := NewColumn(index)
		battery.columnList = append(battery.columnList, *Column)
	}
	return battery
}

// NewColumn is where Elevator are create

func NewColumn(nbEleInColumn int) *Column {
	Column := new(Column)
	Column.nbEleInColumn = 5
	for index := 0; index < Column.nbEleInColumn; index++ {
		Elevator := NewElevator()
		Column.ElevatorList = append(Column.ElevatorList, *Elevator)
	}
	return Column
}

// NewElevator is where Elevator are Define

func NewElevator() *Elevator {
	Elevator := new(Elevator)  
	Elevator.currentFloor = 7
	Elevator.floorList = []int{}
	Elevator.ElevatorStatus = "idle"
	Elevator.ElevatorDirection = "stop"
	Elevator.doorSensor = true
	return Elevator
}

// -------------------- List of all methods--------------------------

// here is the request made by people that want to go down
func (controller *ElevatorController) RequestElevator(FloorNumber, RequestedFloor int) Elevator {
	fmt.Println("Request Elevator to floor : ", FloorNumber)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Button Light On")
	var Column = controller.batteries[0].FindBestColumn(FloorNumber)
	controller.userMove = "down"
	var Elevator = Column.FindBestElevator(RequestedFloor, controller.userMove)
	Elevator.Send_request(FloorNumber)
	Elevator.Send_request(RequestedFloor)
	return Elevator
}

// here is the request from people that want to go up to a floor X
func (controller *ElevatorController) AssignElevator(RequestedFloor int) Elevator {
	fmt.Println("Request Elevator to floor : ", RequestedFloor)
	time.Sleep(3 * time.Millisecond)
	fmt.Println("Button Light On")
	Column := controller.batteries[0].FindBestColumn(RequestedFloor)
	controller.userMove = "up"
	var Elevator = Column.FindBestElevator(RequestedFloor, controller.userMove)
	var FloorNumber = 1
	Elevator.Send_request(FloorNumber)
	Elevator.Send_request(RequestedFloor)
	return Elevator
}

// here is where the best Column are find
func (b *Battery) FindBestColumn(RequestedFloor int) Column { // not sure about *
	if RequestedFloor >= 1 && RequestedFloor <= 7 {
		return b.columnList[0]
	} else if RequestedFloor > 8 && RequestedFloor <= 27 {
		return b.columnList[1]
	} else if RequestedFloor > 28 && RequestedFloor <= 47 {
		return b.columnList[2]
	} else if RequestedFloor > 48 && RequestedFloor <= 66 {
		return b.columnList[3]
	}
	return b.columnList[3]
}

// here is where the best Elevator are found
func (c *Column) FindBestElevator(RequestedFloor int, userMove string) Elevator {
	var selected_Elevator = c.ElevatorList[0]
	for _, e := range c.ElevatorList {  
		if RequestedFloor < e.currentFloor	  && e.ElevatorDirection == "down" && userMove == "down" {
			selected_Elevator = e
		} else if e.ElevatorStatus == "idle" {
			selected_Elevator = e
		} else if e.ElevatorDirection != userMove && e.ElevatorStatus == "moving" || e.ElevatorStatus == "stopped" {
			selected_Elevator = e
		} else if e.ElevatorDirection == userMove && e.ElevatorStatus == "moving" || e.ElevatorStatus == "stopped" {
			selected_Elevator = e
		}
	}
	return selected_Elevator
}

// sendrequest receive information that people made
//in the requestElevator and assign Elevator and sort my list
func (e *Elevator) Send_request(RequestedFloor int) {
	e.floorList = append(e.floorList, RequestedFloor)
	if RequestedFloor > e.currentFloor	  {

		sort.Ints(e.floorList)    
	} else if RequestedFloor < e.currentFloor	  {

		sort.Sort(sort.Reverse(sort.IntSlice(e.floorList)))    
	}
	e.Operate_Elevator(RequestedFloor)
}

// here is where the task are separate depending on the direction
func (e *Elevator) Operate_Elevator(RequestedFloor int) {
	if RequestedFloor == e.currentFloor	  {
		e.OpenDoor()  
	} else if RequestedFloor > e.currentFloor	  {
		e.ElevatorStatus = "moving"
		e.Move_up(RequestedFloor)
		e.ElevatorStatus = "stopped"
		e.OpenDoor()
		e.ElevatorStatus = "moving"  
	} else if RequestedFloor < e.currentFloor	  {
		e.ElevatorStatus = "moving"
		e.Move_down(RequestedFloor)
		e.ElevatorStatus = "stopped"
		e.OpenDoor()
		e.ElevatorStatus = "moving"
	}
}

// here is OpenDoor and CloseDoor
func (e *Elevator) OpenDoor() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Door is Opening")
	time.Sleep(1 * time.Second)
	fmt.Println("Door is Open")
	time.Sleep(1 * time.Second)
	fmt.Println("Button Light Off")
	e.CloseDoor()
}
func (e *Elevator) CloseDoor() {
	if e.doorSensor == true {
		fmt.Println("Door is Closing")
		time.Sleep(1 * time.Second)
		fmt.Println("Door is Close")
		time.Sleep(1 * time.Second)
		fmt.Println("---------------------------------------------------")
		time.Sleep(1 * time.Second)
	} else if e.doorSensor {
		e.OpenDoor()
		fmt.Println("Door cant be close please make sur door is not obstruct")
	}
}

//here is move_up and move_down
func (e *Elevator) Move_up(RequestedFloor int) {  
	fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Current Floor :", e.currentFloor	 )
	for RequestedFloor >   e.currentFloor	  {
		e.currentFloor	  += 1  
		if RequestedFloor == e.currentFloor	  {
			time.Sleep(1 * time.Second)
			fmt.Println("---------------------------------------------------")  
			fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Arrived at destination floor : ", e.currentFloor	 )
		}
		time.Sleep(300 * time.Millisecond)  
		fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Floor : ", e.currentFloor	 )
	}
}
func (e *Elevator) Move_down(RequestedFloor int) {  
	fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Current Floor :", e.currentFloor	 )
	for RequestedFloor <   e.currentFloor	  {
		e.currentFloor	  -= 1  
		if RequestedFloor == e.currentFloor	  {
			time.Sleep(1 * time.Second)
			fmt.Println("---------------------------------------------------")  
			fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Arrived at destination floor : ", e.currentFloor	 )
		}
		time.Sleep(300 * time.Millisecond)  
		fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Floor : ", e.currentFloor	 )
	}
}

func main() {
	controller := NewController(1)

	
	controller.AssignElevator(36)
	controller.RequestElevator(33, 7)
}
