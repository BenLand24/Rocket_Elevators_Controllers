javascript test
remove the comment at line 215 to 222
(those line)

Battery.Column.elelist[0].currentfloor = 2;

// Battery.Column.elelist[1].currentfloor = 6;

// requestelenumber1 = Battery.requestele(7, "down");

// Battery.requestfloor(requestelenumber1, 5);

if you want to add the test 2 in the test 1 remove the comment at line 230 at 238

(those line)

// requestelenumber1 = Battery.requestele(1, "up");

// Battery.requestfloor(requestelenumber1,7);

// requestelenumber2 = Battery.requestele(3, "up");

// Battery.requestfloor(requestelenumber2, 5);

// requestelenumber1 = Battery.requestele(9, "down");

But if you want only test 2 it will be 226 to 238

(those line)
// Battery.Column.elelist[0].currentfloor = 5;

// Battery.Column.elelist[1].currentfloor = 3;

// requestelenumber1 = Battery.requestele(1, "up");

// Battery.requestfloor(requestelenumber1,7);

// requestelenumber2 = Battery.requestele(3, "up");

// Battery.requestfloor(requestelenumber2, 5);

// requestelenumber1 = Battery.requestele(9, "down");

test1

elevator are at floor 2 and floor 6 and someone request at the 7th floor to go down when elevator got it he request to go at floor 5th

test2

elevator are at floor 5 and floor 3 and got requested at floor 1 and floor 3 (so elevator 2 get the request floor 1 and elevator1 will get request 3) floor 1 request the floor 7 and the floor 3 request at floor 5.
and, another request got at floor 9 to go down when elevator get him he go to floor 1  

****For python****

at line 447 to line 452
(those line)
            #first user
            #elevator.append(controller.requestelevator(10, down))
            #second user
            #elevator.append(controller.requestelevator(3, down))
            #third user
            #elevator.append(controller.requestelevator(9, down))
            #first user
            #controller.requestfloor(elevator[0],5)
            #second user
            #controller.requestfloor(elevator[1],2)
            #third user
            #controller.requestfloor(elevator[2],1)

remove the # to unlock the user 

like you want to get first user 

you have to remove the # below first user

second user and third user is the same 
 
and if you want to change the elevator start is the line 429

inittest(10,4) (first start at 9th floor and second start a 3rd floor)

test 

first user request 9th floor to 4th floor, 
second user request 2nd floor to 1st floor,
and third user  request 8th floor to ground

