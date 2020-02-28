using System;
using System.Collections.Generic;
using System.Threading;

namespace Commercial_controller
{
    
    public class ElevatorController
    {
        public int nbFloor;
        public int nbElevatorInColumn;
        public int nbColumn;
        public string userMove;
        public Battery battery;
        public List<int> shortestList;

        public ElevatorController(int nbFloor, int nbColumn, int nbElevatorInColumn, string userMove)
        {
            this.nbFloor = nbFloor;
            this.nbColumn = nbColumn;
            this.nbElevatorInColumn = nbElevatorInColumn;
            this.userMove = userMove;
            this.battery = new Battery(this.nbColumn);
        }


        public Elevator RequestElevator(int floorNumber, int requestedFloor)
        {
            Thread.Sleep(2000);
            Console.WriteLine("Request elevator to floor : " + floorNumber);
            Thread.Sleep(2000);
            var column = battery.findColumn(floorNumber);
            userMove = "down";
            var elevator = column.findRequestedEle(floorNumber, userMove);
            if (requestedFloor <= 0 && requestedFloor > 67)
            {
                Console.WriteLine("this floor is not vailde, try again if it fails again, maybe you're in the wrong building?");
            }
            else
            {
            if (elevator.currentFloor > floorNumber && elevator.Direction == "down") 
            {
                elevator.Request(floorNumber, column.columnNumber);
                elevator.Request(requestedFloor, column.columnNumber);
            }
            else if (elevator.currentFloor < floorNumber && elevator.Direction == "up")
            {
                elevator.Move_down(requestedFloor, column.columnNumber);
                elevator.Request(floorNumber, column.columnNumber);
                elevator.Request(requestedFloor, column.columnNumber);
            }
            Console.WriteLine("Go back to idle");
            }
            return elevator;
        }
    //door/movement
    public class Elevator
    {
        public int eleNumber;
        public string status;
        public int currentFloor;
        public string Direction;
        public bool sensor;
        public int floorDisplay;
        public List<int> floorList;

        public List<string> floorName;
        public Elevator(int eleNumber, string status, int currentFloor, string Direction)
        {
            this.eleNumber = eleNumber;
            this.status = status;
            this.currentFloor = currentFloor;
            this.Direction = Direction;
            this.floorDisplay = currentFloor;
            this.sensor = true;
            this.floorList = new List<int>();
            this.floorName = new List<string>();
        }

        // sendrequest receive information that user requested elevator and assign elevator
        public void Request(int requestedFloor, char columnNumber)
        {
            floorList.Add(requestedFloor);
            if (requestedFloor > currentFloor)
            {
                floorList.Sort((a, b) => a.CompareTo(b));
            }
            else if (requestedFloor < currentFloor)
            {
                floorList.Sort((a, b) => -1 * a.CompareTo(b));

            }

            checkEle(requestedFloor, columnNumber);
        }

        // task for 2 direction (up/down)
        public void checkEle(int requestedFloor, char columnNumber)
        {
            if (requestedFloor == currentFloor)
            {
                Open_door();
                this.status = "idle";

                this.floorList.Remove(0);
            }
            else if (requestedFloor < this.currentFloor)
            {
                status = "moving";
                Console.WriteLine("Column : " + columnNumber + " Elevator : " + this.eleNumber + " " + status);
                this.Direction = "down";
                Move_down(requestedFloor, columnNumber);
                this.status = "stopped";
                Console.WriteLine("Column : " + columnNumber + " Elevator : " + this.eleNumber + " " + status);

                this.Open_door();
                this.floorList.Remove(0);
            }
            else if (requestedFloor > this.currentFloor)
            {
                Thread.Sleep(2000);
                this.status = "moving";
                Console.WriteLine("Column : " + columnNumber + " Elevator : " + this.eleNumber + " " + status);
                this.Direction = "up";
                this.Move_up(requestedFloor, columnNumber);
                this.status = "stopped";
                Console.WriteLine("Column : " + columnNumber + " Elevator : " + this.eleNumber + " " + status);

                this.Open_door();
                this.floorList.Remove(0);
            }

        }
        // Open/Close Door
        public void Open_door()
        {
            Thread.Sleep(2000);

            Console.WriteLine("Door is Opening");

            Thread.Sleep(200);
            Console.WriteLine("Door is Open");
            Thread.Sleep(2000);

            this.Close_door();
        }
        public void Close_door()
        {
            if (sensor == true)
            {
                Console.WriteLine("Door is Closing");
                Thread.Sleep(200);
                Console.WriteLine("Door is Close");
                Thread.Sleep(200);
            }
            else if (sensor == false)
            {
                Open_door();
            }
        }

        // moving direction

        public void Move_up(int requestedFloor, char columnNumber)
        {
            Console.WriteLine("Column : " + columnNumber + " Elevator : #" + eleNumber + "  Current Floor : " + this.currentFloor);
            Thread.Sleep(200);
            while (this.currentFloor != requestedFloor)
            {
                this.currentFloor += 1;
                Console.WriteLine("Column : " + columnNumber + " Elevator : #" + eleNumber + "  Floor : " + this.currentFloor);

                Thread.Sleep(200);
            }
        }

        public void Move_down(int requestedFloor, char columnNumber)
        {
            Console.WriteLine("Column : " + columnNumber + " Elevator : #" + eleNumber + "  Current Floor : " + this.currentFloor);
            Thread.Sleep(200);

            while (this.currentFloor != requestedFloor)
            {
                this.currentFloor -= 1;
                Console.WriteLine("Column : " + columnNumber + " Elevator : #" + eleNumber + "  Floor : " + this.currentFloor);

                Thread.Sleep(200);
            }

        }

    }

    // here is where the column redirect task to find the best column and create my elevator

    public class Column
    {
        public char columnNumber;
        public int nbFloor;
        public int nbElevatorInColumn;
        public List<Elevator> eleList;
        public List<int> call_button_list;


        public Column(char columnNumber, int nbFloor, int nbElevatorInColumn)
        {
            this.columnNumber = columnNumber;
            this.nbFloor = nbFloor;
            this.nbElevatorInColumn = nbElevatorInColumn;
            eleList = new List<Elevator>();
            call_button_list = new List<int>();
            for (int i = 0; i < this.nbElevatorInColumn; i++)
            {
                Elevator elevator = new Elevator(i, "idle", 7, "up");
                eleList.Add(elevator);
            }
        }

        //find best elevator to user request to go up

        public Elevator findAssignElevator(int requestedFloor, int floorNumber, string userMove)
        {

            foreach (var elevator in eleList)
                if (elevator.status == "idle")
                {
                    return elevator;
                }

            var bestElevator = 0;
            var shortest_distance = 1000;
            for (var i = 0; i < this.eleList.Count; i++)
            {
                var ref_distance = Math.Abs(eleList[i].currentFloor - eleList[i].floorList[0]) + Math.Abs(eleList[i].floorList[0] - 1);
                if (shortest_distance >= ref_distance)
                {
                    shortest_distance = ref_distance;
                    bestElevator = i;
                }
            }
            return eleList[bestElevator];
        }

        //find best elevator to user request to go down

        public Elevator findRequestedEle(int requestedFloor, string userMove)
        {
            var shortest_distance = 999;
            var bestElevator = 0;

            for (var i = 0; i < this.eleList.Count; i++)
            {
                var ref_distance = eleList[i].currentFloor - requestedFloor;

                if (ref_distance > 0 && ref_distance < shortest_distance)
                {
                    shortest_distance = ref_distance;
                    bestElevator = i;
                }
            }
            return eleList[bestElevator];
        }

    }
    // here is the battery the things that create my column so that my elevator can work 
    public class Battery
    {
        public string battery_status;
        public int nbColumn;
        public List<Column> column_list;
       

        public Battery(int nbColumn)
        {
            this.nbColumn = nbColumn;
            this.battery_status = "on";
            column_list = new List<Column>();



            char cols = 'A';
            for (int i = 0; i < this.nbColumn; i++, cols++)
            {

                Column column = new Column(cols, 66, 5);

                column.columnNumber = cols;
                column_list.Add(column);

            }
        }
        // here is where the best column are find
        public Column findColumn(int requestedFloor)
        {
            Column best_column = null;
            foreach (Column column in column_list)
            {
                if (requestedFloor > 1 && requestedFloor <= 7 || requestedFloor == 7)
                {
                    best_column = column_list[0];
                }
                else if (requestedFloor > 8 && requestedFloor <= 27 || requestedFloor == 7)
                {

                    best_column = column_list[1];


                }
                else if (requestedFloor > 28 && requestedFloor <= 47 || requestedFloor == 7)
                {
                    best_column = column_list[2];


                }
                else if (requestedFloor > 48 && requestedFloor <= 66 || requestedFloor == 7)
                {
                    best_column = column_list[3];


                }
            }
            return best_column;
        }
    }
        //request user how want go up at floor X
        public Elevator assignElevator(int requestedFloor)
        {
            Thread.Sleep(2000);
            Console.WriteLine("Requested floor : " + requestedFloor);
            Thread.Sleep(2000);
            Console.WriteLine("Call Button Light On");


            Column column = battery.findColumn(requestedFloor);
            var floorNumber = 7;
            Elevator elevator = column.findAssignElevator(requestedFloor, floorNumber, userMove);

            elevator.Request(floorNumber, column.columnNumber);

            elevator.Request(requestedFloor, column.columnNumber);

            return elevator;
        }

    


    public class CommercialCS
    {
        public static void Main(string[] args)
        {
            ElevatorController controller = new ElevatorController(66, 4, 5, "stop");


        //Scenario

                           //Méthode 1: RequestElevator(floorNumber, requestedFloor)
                          //Méthode 2: assignElevator(requestedFloor)
                         //floorNumber = the place where the customer is at
                        //requestedFloor = the floor the customer want to go


        //Column A

            controller.battery.column_list[0].eleList[0].currentFloor = 3;
            controller.battery.column_list[0].eleList[0].Direction = "stop";
            controller.battery.column_list[0].eleList[0].status = "idle";
            controller.battery.column_list[0].eleList[0].floorList.Add(3);


            controller.battery.column_list[0].eleList[1].currentFloor = 7;
            controller.battery.column_list[0].eleList[1].Direction = "stop";
            controller.battery.column_list[0].eleList[1].status = "idle";
            controller.battery.column_list[0].eleList[1].floorList.Add(7);


            controller.battery.column_list[0].eleList[2].currentFloor = 4;
            controller.battery.column_list[0].eleList[2].Direction = "down";
            controller.battery.column_list[0].eleList[2].status = "moving";
            controller.battery.column_list[0].eleList[2].floorList.Add(2);


            controller.battery.column_list[0].eleList[3].currentFloor = 1;
            controller.battery.column_list[0].eleList[3].Direction = "up";
            controller.battery.column_list[0].eleList[3].status = "moving";
            controller.battery.column_list[0].eleList[3].floorList.Add(7);


            controller.battery.column_list[0].eleList[4].currentFloor = 6;
            controller.battery.column_list[0].eleList[4].Direction = "down";
            controller.battery.column_list[0].eleList[4].status = "moving";
            controller.battery.column_list[0].eleList[4].floorList.Add(1);
           
            Elevator elevator = controller.RequestElevator(4, 7);
        //column B

            // controller.battery.column_list[1].eleList[0].currentFloor = 25;
            // controller.battery.column_list[1].eleList[0].Direction = "down";
            // controller.battery.column_list[1].eleList[0].status = "moving";
            // controller.battery.column_list[1].eleList[0].floorList.Add(12);


            // controller.battery.column_list[1].eleList[1].currentFloor = 10;
            // controller.battery.column_list[1].eleList[1].Direction = "up";
            // controller.battery.column_list[1].eleList[1].status = "moving";
            // controller.battery.column_list[1].eleList[1].floorList.Add(13);


            // controller.battery.column_list[1].eleList[2].currentFloor = 20;
            // controller.battery.column_list[1].eleList[2].Direction = "down";
            // controller.battery.column_list[1].eleList[2].status = "moving";
            // controller.battery.column_list[1].eleList[2].floorList.Add(7);


            // controller.battery.column_list[1].eleList[3].currentFloor = 22;
            // controller.battery.column_list[1].eleList[3].Direction = "down";
            // controller.battery.column_list[1].eleList[3].status = "moving";
            // controller.battery.column_list[1].eleList[3].floorList.Add(9);


            // controller.battery.column_list[1].eleList[4].currentFloor = 13;
            // controller.battery.column_list[1].eleList[4].Direction = "down";
            // controller.battery.column_list[1].eleList[4].status = "moving";
            // controller.battery.column_list[1].eleList[4].floorList.Add(7);

            // controller.assignElevator(27);
            // Elevator elevator = controller.RequestElevator(7, 27);

        //Column C
            // controller.battery.column_list[2].eleList[0].currentFloor = 7;
            // controller.battery.column_list[2].eleList[0].Direction = "up";
            // controller.battery.column_list[2].eleList[0].status = "moving";
            // controller.battery.column_list[2].eleList[0].floorList.Add(28);


            // controller.battery.column_list[2].eleList[1].currentFloor = 30;
            // controller.battery.column_list[2].eleList[1].Direction = "up";
            // controller.battery.column_list[2].eleList[1].status = "moving";
            // controller.battery.column_list[2].eleList[1].floorList.Add(35);


            // controller.battery.column_list[2].eleList[2].currentFloor = 47;
            // controller.battery.column_list[2].eleList[2].Direction = "down";
            // controller.battery.column_list[2].eleList[2].status = "moving";
            // controller.battery.column_list[2].eleList[2].floorList.Add(31);


            // controller.battery.column_list[2].eleList[3].currentFloor = 46;
            // controller.battery.column_list[2].eleList[3].Direction = "down";
            // controller.battery.column_list[2].eleList[3].status = "moving";
            // controller.battery.column_list[2].eleList[3].floorList.Add(7);


            // controller.battery.column_list[2].eleList[4].currentFloor = 47;
            // controller.battery.column_list[2].eleList[4].Direction = "down";
            // controller.battery.column_list[2].eleList[4].status = "moving";
            // controller.battery.column_list[2].eleList[4].floorList.Add(31);

            // controller.assignElevator(43);
            // Elevator elevator = controller.RequestElevator(7, 43);


        //column D
            // controller.battery.column_list[3].eleList[0].currentFloor = 64;
            // controller.battery.column_list[3].eleList[0].Direction = "down";
            // controller.battery.column_list[3].eleList[0].status = "moving";
            // controller.battery.column_list[3].eleList[0].floorList.Add(7);


            // controller.battery.column_list[3].eleList[1].currentFloor = 57;
            // controller.battery.column_list[3].eleList[1].Direction = "up";
            // controller.battery.column_list[3].eleList[1].status = "moving";
            // controller.battery.column_list[3].eleList[1].floorList.Add(66);


            // controller.battery.column_list[3].eleList[2].currentFloor = 53;
            // controller.battery.column_list[3].eleList[2].Direction = "up";
            // controller.battery.column_list[3].eleList[2].status = "moving";
            // controller.battery.column_list[3].eleList[2].floorList.Add(64);


            // controller.battery.column_list[3].eleList[3].currentFloor = 7;
            // controller.battery.column_list[3].eleList[3].Direction = "up";
            // controller.battery.column_list[3].eleList[3].status = "moving";
            // controller.battery.column_list[3].eleList[3].floorList.Add(60);


            // controller.battery.column_list[3].eleList[4].currentFloor = 66;
            // controller.battery.column_list[3].eleList[4].Direction = "down";
            // controller.battery.column_list[3].eleList[4].status = "moving";
            // controller.battery.column_list[3].eleList[4].floorList.Add(7);

            // Elevator elevator = controller.RequestElevator(54, 7);

            }
        }
    }
}