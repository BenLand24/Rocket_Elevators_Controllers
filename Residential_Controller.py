
import itertools
import os
import time
import datetime
from datetime import timedelta

nbcolumn = 1
nbelev = 2
nbfloors = 10
floornames = ["Ground", "1st", "2nd", "3rd","4th", "5th", "6th", "7th", "8th", "9th"]
Ground = 1
defaultfloor = 3
timeperfloor = 1000

#elevator directions
directionnamelist = ["usermove", "down", "up"]
usermove = 0
down = 1
up = 2

#people entering or leaving the elevator
usermovenamelist = ["movein","moveout"]
movein = 0
moveout = 1

#elevator status
elevatorstatusnamelist = ["idle","stopping","stopped","moving"]
idle = 0
stopping = 1
stopped = 2
moving = 3

#door status
doorstatusnamelist = ["closed","closing","opening","opened"]
closed = 0
closing = 1
opening = 2
opened = 3

#buttons function
addfloor = 0
callelevator = 1
opendoor = 2
closedoor = 3

#buttons status
inactive = 0
active = 1

delayelevatorstopping = 2000
delaydooropening = 1000
delaybeforeclosedoor = 5000
delayforceclosedoor = 10000
timeoutdooropen = 7500
delaymaxidletime = 7500   
currentfloor = 0
apptimeout = 60000

timeinmilli = lambda: int(round(time.time() * 1000))

def counter(): #Self increment counter to generate unique ID's
    return lambda c=itertools.count(): next(c)

count = counter()  


class elevatorcontroller:
    def __init__(self):
        self.id = 1
        self.columns = column(2, Ground, 2)
        print("controller Created a column for " + str(nbelev) + " elevators and for " + str(nbfloors) + " floors")

    def requestelevator(self, floornumber, direction):
        print("elevator requested at " + floornames[floornumber-1] + " to go " + directionnamelist[direction])
        elevator = self.callelevator(direction, floornumber, movein)
        return elevator

    def requestfloor(self, elevator, requestedfloor):
        self.adddestinationelev(requestedfloor, elevator, moveout,"")

    def callelevator(self, requesteddirection, requestedfloor, usermove):
        elevator = self.findelevator(requesteddirection, requestedfloor)
        if self.checkifdestinexist(requestedfloor, elevator) is False:
            self.adddestinationelev(requestedfloor, elevator, usermove, requesteddirection)
        return elevator

    def checkifdestinexist(self, floor, elevator):
        if elevator.destinationlist:
            for destination in elevator.destinationlist:
                if destination == floor:
                    return True
        return False

    def findelevator(self, requesteddirection, requestedfloor):
        idleelevatorlist = []
        for elevator in self.columns.elevators:
            if elevator.door.alarm is False:
                if requestedfloor is elevator.currentfloor:
                    if (elevator.status is stopped and elevator.direction is requesteddirection) or (elevator.status is idle and not elevator.destinationlist):
                        return elevator
                elif ((requestedfloor >= elevator.currentfloor) and ((elevator.direction is up) or (elevator.direction is usermove)) and (requesteddirection is up) and ((elevator.status is moving) or (elevator.status is stopped))):
                    return elevator
                elif ((requestedfloor <= elevator.currentfloor) and ((elevator.direction is down) or (elevator.direction is usermove)) and (requesteddirection is down) and ((elevator.status is moving) or (elevator.status is stopped))):
                    return elevator
                elif elevator.status is idle and not elevator.destinationlist:
                    idleelevatorlist.append(elevator)
        if idleelevatorlist:
            if (len(idleelevatorlist) > 1):
                gap = 999
                for elevator in idleelevatorlist:
                    if gap > abs(elevator.currentfloor - requestedfloor):
                        gap = abs(elevator.currentfloor - requestedfloor)
                        elevatortouse = elevator
                return elevatortouse
            else:
                return idleelevatorlist[0]

        elevator = self.nearestelevator(requestedfloor, requesteddirection)
        if not elevator:
            return elevator
        else:
            return self.shortestdestinationlist()

    def shortestdestinationlist(self):
        length = 999
        for elevator in self.columns.elevators:
            if length > len(elevator.destinationlist):
                length = len(elevator.destinationlist)
                elevwithshortestlist = elevator
        return elevwithshortestlist
    
    def nearestelevator(self, requestedfloor, requesteddirection):
        gap = 999
        elevwithshortestgap = None
        for elevator in self.columns.elevators:
            if (gap > abs(elevator.currentfloor - elevator.destinationlist[0])) and (elevator.door.alarm is False):
                if ((requestedfloor > elevator.currentfloor) and ((elevator.direction is up) or (elevator.direction is usermove)) and (requesteddirection is up)) or ((requestedfloor > elevator.currentfloor) and ((elevator.direction is down) or (elevator.direction is usermove)) and (requesteddirection is down)):
                    gap = abs(elevator.currentfloor - elevator.destinationlist[0])
                    elevwithshortestgap = elevator
        return elevwithshortestgap

    def adddestinationelev(self, floor, elevator, usermove, requesteddirection):
        if elevator.destinationlist:
            for destination in elevator.destinationlist:
                if elevator.direction:
                    if elevator.direction is up:
                        if (floor < destination) and (floor > elevator.currentfloor):
                            elevator.destinationlist.sort()
                            return
                        elif (floor > destination) and (floor < elevator.currentfloor):
                            elevator.destinationlist.sort()
                            return
                    elif elevator.direction is down:
                        if (floor > destination) and (floor < elevator.currentfloor):
                            elevator.destinationlist.sort(reversed)
                            return
                        elif (floor < destination) and (floor > elevator.currentfloor):
                            elevator.destinationlist.sort(reversed)
                            return
                else:
                    elevator.destinationlist.append(floor)
                    elevator.userlist.append(usermove)
                    elevator.directionlist.append(requesteddirection)
                    return
        else:
            elevator.destinationlist.append(floor)
            elevator.userlist.append(usermove)
            elevator.directionlist.append(requesteddirection)
            return

    def checkelevatorstatus(self):
        for elevator in self.columns.elevators:
            if (timeinmilli() > elevator.opendoortime + timeoutdooropen) and (elevator.door.status is not closed):
                elevator.door.alarm = True
                elevator.forceclosedoor()
            if (timeinmilli() > elevator.idletime + delaymaxidletime) and (not elevator.destinationlist):
                pass

            if (timeinmilli() > (elevator.opendoortime + delaydooropening)) and (elevator.door.status is opening):
                elevator.door.status = opened
                print("elevator " + str(elevator.id) + " door is opened") 
            
            if ((timeinmilli() > (elevator.opendoortime + delaybeforeclosedoor)) and (elevator.door.status is opened)) or elevator.door.alarm is True:
                elevator.closedoor()
            
            if (timeinmilli() > elevator.closedoortime + delaydooropening) and elevator.door.status is closing:
                if (elevator.door.alarm is True) and (elevator.door.status is not opening):
                    elevator.opendoor()
                    print("elevator " + str(elevator.id) + " door is opening again because of obstruction")
                elevator.door.status = closed
                elevator.status = idle
                elevator.idletime = timeinmilli()
                print("elevator " + str(elevator.id) + " door is closed")
            
            if (timeinmilli() > elevator.forceclosetime + delayforceclosedoor) and (elevator.door.status is closing) and elevator.door.alarm is False:
                elevator.status = idle
                elevator.idletime = timeinmilli()
                elevator.alarm = False
                elevator.door.alarm = False
                elevator.door.status = closed
                print("elevator " + str(elevator.id) + " door is closed")

            if (timeinmilli() > elevator.stoptime + delayelevatorstopping) and elevator.status is stopping:
                elevator.status = stopped
                elevator.direction = usermove
                print("elevator " + str(elevator.id) + " is stopped at " + str(floornames[elevator.currentfloor-1]) + " floor for people " + usermovenamelist[elevator.userlist[0]])
                elevator.directionlist.pop(0)
                elevator.userlist.pop(0)

    def verifydestinationlist(self):
        for elevator in self.columns.elevators:
            if elevator.destinationlist:
                if elevator.userlist:
                    if (elevator.destinationlist[0]) and (elevator.status is idle):
                        if (elevator.currentfloor != elevator.destinationlist[0]) and elevator.door.status is closed:
                            elevator.startmove()
                        elif (elevator.userlist[0] is movein) and (elevator.currentfloor is elevator.destinationlist[0]) and elevator.door.status is closed:
                            #print("open door 1")
                            elevator.destinationlist.pop(0)
                            elevator.directionlist.pop(0)
                            elevator.userlist.pop(0)
                            elevator.opendoor()

    def checkmovingelevator(self):
        for elevator in self.columns.elevators:
            if (elevator.status is moving) and (elevator.door.alarm is False):
                if (timeinmilli() - elevator.movetimestamp) >= timeperfloor:
                    if elevator.currentfloor < elevator.destinationlist[0]:
                        elevator.floorlevelfortiming += 1
                    else:
                        elevator.floorlevelfortiming -= 1

                    elevator.movetimestamp = timeinmilli()
                    if elevator.currentfloor is not elevator.floorlevelfortiming:
                        elevator.currentfloor = elevator.floorlevelfortiming
                        print("elevator " + str(elevator.id) + " is at " + floornames[elevator.currentfloor-1] + " floor")
                        
                    if elevator.currentfloor is elevator.destinationlist[0]:
                        elevator.stopelevator()
                        self.clearbuttons(elevator)
                        elevator.destinationlist.pop(0)
                        print(elevator.destinationlist)

            if elevator.status is stopped and elevator.door.status is closed:
                elevator.opendoor()

    def clearbuttons(self, elevator):
        if elevator.directionlist[0] != "":
            print(elevator.directionlist[0], elevator.destinationlist[0])
            button = self.columns.finddirectionbutton(elevator.directionlist[0],elevator.destinationlist[0]) 
            button.status = inactive
            print("request button direction " + directionnamelist[button.direction] + " at " + floornames[elevator.destinationlist[0]-1] + " floor is inactive")
        button = elevator.findfloorbutton(elevator.destinationlist[0])
        button.status = inactive
        print("elevator " + str(elevator.id) + " " + floornames[elevator.currentfloor-1] + " floor button is inactive")


class floor:
    def __init__(self, id, name, buttons):
        self.id = id
        self.name = name



class column:
    def __init__(self, nbElev, defaultFloor, nbdirectionbuttons):
        self.nbelev = nbelev
        self.defaultfloor = defaultfloor
        self.nbdirectionbuttons = nbdirectionbuttons
        self.elevators = []
        for index in range(self.nbelev):
            self.elevators.append(elevator(index+1))
        self.destinationlist = []
        self.directionbuttonlist = []
        for index in range(nbfloors):
            if index == 0:
                self.directionbuttonlist.append(directionbutton(up, index))
            elif index == (nbfloors - 1):
                self.directionbuttonlist.append(directionbutton(down, index))
            else:
                self.directionbuttonlist.append(directionbutton(up, index))

                self.directionbuttonlist.append(directionbutton(down, index))

            self.destinationlist.append(floor(index, floornames[index], self.directionbuttonlist))

    def finddirectionbutton(self, requesteddirection, requestedfloor):
        for button in self.directionbuttonlist:
            if (button.direction is requesteddirection) and (button.floor is requestedfloor):
                return button


class elevator:
    def __init__(self, id):
        self.id = id
        self.currentfloor = Ground
        self.direction = usermove
        self.status = idle
        self.floorlevelfortiming = Ground
        self.alarm = False
        self.movetimestamp = 0
        self.idletime = timeinmilli()
        self.opendoortime = timeinmilli()
        self.closedoortime = timeinmilli()
        self.stoptime = timeinmilli()
        self.forceclosetime = timeinmilli()
        self.destinationlist = []
        self.userlist = []
        self.directionlist = []
        self.floorbuttons = []  #buttons inside the elevators
        for index in range(nbfloors):
            self.floorbuttons.append(floorbutton(index))

        self.door = door()
        self.opendoorbutton = opendoorbutton(self.door)
        self.closedoorbutton = closedoorbutton(self.door)
        print("elevator created")

    def findfloorbutton (self, currentfloor):
        for button in self.floorbuttons: 
            if button.floor is self.currentfloor:
                return button 

    def startmove(self):
        if (self.currentfloor < self.destinationlist[0]): 
            self.direction = up
        elif (self.currentfloor > self.destinationlist[0]): 
            self.direction = down
        self.status = moving
        self.movetimestamp = timeinmilli()
        self.floorlevelfortiming = self.currentfloor
        print("elevator " + str(self.id) + " is moving " + directionnamelist[self.direction] + " to " + floornames[self.destinationlist[0]-1] + " floor for people " + usermovenamelist[self.userlist[0]])

    def stopelevator(self):
        self.stoptime = timeinmilli()
        self.status = stopping
        print("elevator " + str(self.id) + " is stopping at " + floornames[self.destinationlist[0]-1] + " floor")

    def opendoor(self):
        self.door.status = opening
        self.opendoortime = timeinmilli()
        print("elevator " + str(self.id) + " door is opening")

    def closedoor(self):
        self.door.status = closing
        self.closedoortime = timeinmilli()
        print("elevator " + str(self.id) + " door is closing")

    def forceclosedoor(self):
        self.alarm = True 
        print("elevator " + str(self.id) + " door is closing slowly (force close)")
        self.forceclosetime = timeinmilli()


class directionbutton:
    def __init__(self, direction, floor):
        self.floor = floor
        self.direction = direction  # up or down
        self.status = inactive  # active or inactive


class floorbutton:
    def __init__(self, floor):
        self.floor = floor
        self.status = inactive


class door:
    def __init__(self):
        self.id = id
        self.status = closed
        self.obstruction = False
        self.alarm = False


class opendoorbutton:
    def __init__(self, door):
        self.door = door


class closedoorbutton:
    def __init__(self, door):
        self.door = door


controller = elevatorcontroller()

def inittest(elev1floor,elev2floor):
    controller.columns.elevators[0].currentfloor = elev1floor
    controller.columns.elevators[0].floorlevelfortiming = elev1floor
    print("elevator 1 current floor is " + floornames[controller.columns.elevators[0].currentfloor-1] + " floor")
    
    controller.columns.elevators[1].currentfloor = elev2floor
    controller.columns.elevators[1].floorlevelFortiming = elev2floor
    print("elevator 2 current floor is " + floornames[controller.columns.elevators[1].currentfloor-1] + " floor")


"""
floor index:
    Ground : 1
    1st : 2
    2nd : 3
    3rd : 4
    4th : 5
    5th : 6
    6th : 7
    7th : 8
    8th : 9
    9th : 10

ELEVATOR INDEX:
    1 : 1
    2 : 2

STATUS INDEX:
    usermove : 0
    down : 1
    up : 2

"""

#controller initialisation
inittest(10,4)
controllerrunning = True

def main():
    elevator = []
    firstinstructiondone = False
    timeout = timeinmilli()
    readytostopcontroller = False 

    while controllerrunning is True:
        #INITIATE ALL THE CALLS
        if firstinstructiondone is False:
            firstinstructiondone = True
            #first user
            elevator.append(controller.requestelevator(10, down))
            #second user
            #elevator.append(controller.requestelevator(3, down))
            #third user
            #elevator.append(controller.requestelevator(9, down))
            #first user
            controller.requestfloor(elevator[0],5)
            #second user
            #controller.requestfloor(elevator[1],2)
            #third user
            #controller.requestfloor(elevator[2],1)

        #CONTROLLER SEQUENCES
        controller.checkelevatorstatus()
        controller.verifydestinationlist()
        controller.checkmovingelevator()

        #STOP THE CONTROLLER WHEN DONE OR TIMED OUT
        if ((timeinmilli() > controller.columns.elevators[0].idletime + apptimeout) and (timeinmilli() > controller.columns.elevators[1].idletime + apptimeout)):
            readytostopcontroller = True
        if (timeinmilli() > timeout + apptimeout) or (readytostopcontroller is True):
            print("break")
            break

main()
