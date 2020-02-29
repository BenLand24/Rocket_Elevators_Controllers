package main

import (
	"fmt"
	"sort"
	"time"
)

// ElevatorController var
type ElevatorController struct {
	BatteryNumber int
	batteries     []Battery
	nbColumn      int
	userMove      string
}

// Battery var
type Battery struct {
	nbColumn   int
	columnList []Column
}

// Column var
type Column struct {
	ColumnNumber  int
	nbEleInColumn int
	EleList  []Elevator
}

// Elevator var
type Elevator struct {
	eleNumber         int
	currentFloor      int
	floorList         []int
	Status    		  string
	Direction 		  string
	doorSensor        bool
	Column            Column
}

// NewController = Battery create

func NewController(BatteryNumber int) ElevatorController {
	controller := new(ElevatorController)
	controller.BatteryNumber = 1
	for index := 0; index < BatteryNumber; index++ {
		battery := NewBattery(index)
		controller.batteries = append(controller.batteries, *battery)
	}
	return *controller
}

// NewBattery = Column create

func NewBattery(nbColumn int) *Battery {
	battery := new(Battery)
	battery.nbColumn = 4
	for index := 0; index < battery.nbColumn; index++ {
		Column := NewColumn(index)
		battery.columnList = append(battery.columnList, *Column)
	}
	return battery
}

// NewColumn = Elevator create

func NewColumn(nbEleInColumn int) *Column {
	Column := new(Column)
	Column.nbEleInColumn = 5
	for index := 0; index < Column.nbEleInColumn; index++ {
		Elevator := NewElevator()
		Column.EleList = append(Column.EleList, *Elevator)
	}
	return Column
}

// NewElevator = Elevator Define

func NewElevator() *Elevator {
	Elevator := new(Elevator)
	Elevator.currentFloor = 7
	Elevator.floorList = []int{}
	Elevator.Status = "idle"
	Elevator.Direction = "stop"
	Elevator.doorSensor = true
	return Elevator
}

// request by user who want to go down
func (controller *ElevatorController) RequestElevator(FloorNumber, RequestedFloor int) Elevator {
	fmt.Println("Request Elevator to floor : ", FloorNumber)
	time.Sleep(2000 * time.Millisecond)
	var Column = controller.batteries[0].FindBestColumn(FloorNumber)
	controller.userMove = "down"
	var Elevator = Column.FindBestElevator(RequestedFloor, controller.userMove)
	Elevator.SendRequest(FloorNumber)
	Elevator.SendRequest(RequestedFloor)
	return Elevator
}

// request by user who want to go up at floor X
func (controller *ElevatorController) AssignElevator(RequestedFloor int) Elevator {
	fmt.Println("Request Elevator to floor : ", RequestedFloor)
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Button Light On")
	Column := controller.batteries[0].FindBestColumn(RequestedFloor) //need to change 0 for ""
	controller.userMove = "up"
	var Elevator = Column.FindBestElevator(RequestedFloor, controller.userMove)
	var FloorNumber = 7
	Elevator.SendRequest(FloorNumber)
	Elevator.SendRequest(RequestedFloor)
	return Elevator
}

// find best Column 
func (b *Battery) FindBestColumn(RequestedFloor int) Column { // not sure 
	if RequestedFloor >= 1 && RequestedFloor <= 7 {
		return b.columnList[0]
	} else if RequestedFloor > 8 && RequestedFloor <= 27 || RequestedFloor == 7 {
		return b.columnList[1]
	} else if RequestedFloor > 28 && RequestedFloor <= 47 || RequestedFloor == 7 {
		return b.columnList[2]
	} else if RequestedFloor > 48 && RequestedFloor <= 66 || RequestedFloor == 7 {
		return b.columnList[3]
	}
	return b.columnList[3] //need to change 3 for ""
}

// find best Elevator
func (c *Column) FindBestElevator(RequestedFloor int, userMove string) Elevator {
	var selectedElevator = c.EleList[4] // need to change 4 for ""(but 5 crash console)
	for _, e := range c.EleList {
		if RequestedFloor < e.currentFloor && e.Direction == "down" && userMove == "down" {
			selectedElevator = e
		} else if e.Status == "idle" {
			selectedElevator = e
		} else if e.Direction != userMove && e.Status == "moving" || e.Status == "stopped" {
			selectedElevator = e
		} else if e.Direction == userMove && e.Status == "moving" || e.Status == "stopped" {
			selectedElevator = e
		}
	}
	return selectedElevator
}

//requestElevator assign Elevator sortlist
func (e *Elevator) SendRequest(RequestedFloor int) {
	e.floorList = append(e.floorList, RequestedFloor)
	if RequestedFloor > e.currentFloor {

		sort.Ints(e.floorList)
	} else if RequestedFloor < e.currentFloor {

		sort.Sort(sort.Reverse(sort.IntSlice(e.floorList)))
	}
	e.OperateElevator(RequestedFloor)
}

// userrequest/move depending on the direction
func (e *Elevator) OperateElevator(RequestedFloor int) {
	if RequestedFloor == e.currentFloor {
		e.OpenDoor()
	} else if RequestedFloor > e.currentFloor {
		e.Status = "moving"
		e.MoveUp(RequestedFloor)
		e.Status = "stopped"
		e.OpenDoor()
		e.Status = "moving"
	} else if RequestedFloor < e.currentFloor {
		e.Status = "moving"
		e.MoveDown(RequestedFloor)
		e.Status = "stopped"
		e.OpenDoor()
		e.Status = "moving"
	}
}

// OpenDoor/CloseDoor
func (e *Elevator) OpenDoor() {
	fmt.Println("Door is Opening")
	time.Sleep(1 * time.Second)
	fmt.Println("Door is Open")
	time.Sleep(1 * time.Second)
	e.CloseDoor()
}
func (e *Elevator) CloseDoor() {
	if e.doorSensor == true {
		fmt.Println("Door is Closing")
		time.Sleep(1 * time.Second)
		fmt.Println("Door is Close")
		time.Sleep(1 * time.Second)
	} else if e.doorSensor {
		e.OpenDoor()
		fmt.Println("Door can not be close please make sur door is not obstruct")
	}
}

// moveUp/moveDown
func (e *Elevator) MoveUp(RequestedFloor int) {
	fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Current Floor :", e.currentFloor)
	for RequestedFloor > e.currentFloor {
		e.currentFloor += 1
		if RequestedFloor == e.currentFloor {
			time.Sleep(1 * time.Second)
			fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Reach the destination floor : ", e.currentFloor)
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Floor : ", e.currentFloor)
	}
}
func (e *Elevator) MoveDown(RequestedFloor int) {
	fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Current Floor :", e.currentFloor)
	for RequestedFloor < e.currentFloor {
		e.currentFloor -= 1
		if RequestedFloor == e.currentFloor {
			time.Sleep(1 * time.Second)
			fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Reach the destination floor : ", e.currentFloor)
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.eleNumber, " Floor : ", e.currentFloor)
	}
}

func main() {
	controller := NewController(1)

	//column A

	// controller.batteries[0].columnList[0].EleList[0].currentFloor = 3
	// controller.batteries[0].columnList[0].EleList[0].Status = "idle"
	// controller.batteries[0].columnList[0].EleList[0].Direction = "stop"
	// // controller.batteries[0].columnList[0].EleList[0].SendRequest(3)

	// controller.batteries[0].columnList[0].EleList[1].currentFloor = 7
	// controller.batteries[0].columnList[0].EleList[1].Status = "idle"
	// controller.batteries[0].columnList[0].EleList[1].Direction = "stop"
	// // controller.batteries[0].columnList[0].EleList[1].SendRequest(7)

	// controller.batteries[0].columnList[0].EleList[2].currentFloor = 4
	// controller.batteries[0].columnList[0].EleList[2].Status = "moving"
	// controller.batteries[0].columnList[0].EleList[2].Direction = "down"
	// //controller.batteries[0].columnList[0].EleList[2].SendRequest(2)

	// controller.batteries[0].columnList[0].EleList[3].currentFloor = 1
	// controller.batteries[0].columnList[0].EleList[3].Status = "moving"
	// controller.batteries[0].columnList[0].EleList[3].Direction = "up"
	// //controller.batteries[0].columnList[0].EleList[3].SendRequest(7)

	// controller.batteries[0].columnList[0].EleList[4].currentFloor = 6
	// controller.batteries[0].columnList[0].EleList[4].Status = "moving"
	// controller.batteries[0].columnList[0].EleList[4].Direction = "down"
	// //controller.batteries[0].columnList[0].EleList[4].SendRequest(1)

	// // controller.AssignElevator(1)
	// controller.RequestElevator(4, 7)

	//column B

	controller.batteries[0].columnList[1].EleList[0].currentFloor = 25
	controller.batteries[0].columnList[1].EleList[0].Status = "moving"
	controller.batteries[0].columnList[1].EleList[0].Direction = "down"
	//controller.batteries[0].columnList[1].EleList[0].SendRequest(12)

	controller.batteries[0].columnList[1].EleList[1].currentFloor = 10
	controller.batteries[0].columnList[1].EleList[1].Status = "moving"
	controller.batteries[0].columnList[1].EleList[1].Direction = "up"
	//controller.batteries[0].columnList[1].EleList[1].SendRequest(13)

	controller.batteries[0].columnList[1].EleList[2].currentFloor = 20
	controller.batteries[0].columnList[1].EleList[2].Status = "moving"
	controller.batteries[0].columnList[1].EleList[2].Direction = "down"
	//controller.batteries[0].columnList[1].EleList[2].SendRequest(7)

	controller.batteries[0].columnList[1].EleList[3].currentFloor = 22
	controller.batteries[0].columnList[1].EleList[3].Status = "moving"
	controller.batteries[0].columnList[1].EleList[3].Direction = "down"
	//controller.batteries[0].columnList[1].EleList[3].SendRequest(9)

	controller.batteries[0].columnList[1].EleList[4].currentFloor = 13
	controller.batteries[0].columnList[1].EleList[4].Status = "moving"
	controller.batteries[0].columnList[1].EleList[4].Direction = "down"
	// controller.batteries[0].columnList[1].EleList[4].SendRequest(7)

	controller.AssignElevator(27)
	// controller.RequestElevator(7, 27)

	//column C

	// controller.batteries[0].columnList[2].EleList[0].currentFloor = 7
	// controller.batteries[0].columnList[2].EleList[0].Status = "moving"
	// controller.batteries[0].columnList[2].EleList[0].Direction = "up"
	// // controller.batteries[0].columnList[2].EleList[0].SendRequest(28)

	// controller.batteries[0].columnList[2].EleList[1].currentFloor = 30
	// controller.batteries[0].columnList[2].EleList[1].Status = "moving"
	// controller.batteries[0].columnList[2].EleList[1].Direction = "up"
	// // controller.batteries[0].columnList[2].EleList[1].SendRequest(35)

	// controller.batteries[0].columnList[2].EleList[2].currentFloor = 47
	// controller.batteries[0].columnList[2].EleList[2].Status = "moving"
	// controller.batteries[0].columnList[2].EleList[2].Direction = "down"
	// // controller.batteries[0].columnList[2].EleList[2].SendRequest(31)

	// controller.batteries[0].columnList[2].EleList[3].currentFloor = 46
	// controller.batteries[0].columnList[2].EleList[3].Status = "moving"
	// controller.batteries[0].columnList[2].EleList[3].Direction = "down"
	// // controller.batteries[0].columnList[2].EleList[3].SendRequest(7)

	// controller.batteries[0].columnList[2].EleList[4].currentFloor = 47
	// controller.batteries[0].columnList[2].EleList[4].Status = "moving"
	// controller.batteries[0].columnList[2].EleList[4].Direction = "down"
	// // controller.batteries[0].columnList[2].EleList[4].SendRequest(31)

	// controller.AssignElevator(43)
	// // controller.RequestElevator(7, 43)

	//column D

	// controller.batteries[0].columnList[3].EleList[0].currentFloor = 64
	// controller.batteries[0].columnList[3].EleList[0].Status = "moving"
	// controller.batteries[0].columnList[3].EleList[0].Direction = "down"
	// // controller.batteries[0].columnList[3].EleList[0].SendRequest(7)

	// controller.batteries[0].columnList[3].EleList[1].currentFloor = 57
	// controller.batteries[0].columnList[3].EleList[1].Status = "moving"
	// controller.batteries[0].columnList[3].EleList[1].Direction = "down"
	// // controller.batteries[0].columnList[3].EleList[1].SendRequest(66)

	// controller.batteries[0].columnList[3].EleList[2].currentFloor = 62
	// controller.batteries[0].columnList[3].EleList[2].Status = "moving"
	// controller.batteries[0].columnList[3].EleList[2].Direction = "down"
	// // controller.batteries[0].columnList[3].EleList[2].SendRequest(64)

	// controller.batteries[0].columnList[3].EleList[3].currentFloor = 7
	// controller.batteries[0].columnList[3].EleList[3].Status = "moving"
	// controller.batteries[0].columnList[3].EleList[3].Direction = "down"
	// // controller.batteries[0].columnList[3].EleList[3].SendRequest(60)

	// controller.batteries[0].columnList[3].EleList[4].currentFloor = 66
	// controller.batteries[0].columnList[3].EleList[4].Status = "moving"
	// controller.batteries[0].columnList[3].EleList[4].Direction = "down"
	// // controller.batteries[0].columnList[3].EleList[4].SendRequest(7)

	// // controller.AssignElevator(60)
	// controller.RequestElevator(54, 7)
}
