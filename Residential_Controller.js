
var nbfloors = 10;
var nbcolumn = 1;
var nbelev = 2;
var nbbattery = 1;
var statusele = this.statusele

function opendoor(){
    if ("statusele = idle" || "statusele = stop"){
        door = open;
    }
    else{
        door = close;
    }
}
function closedoor(){
    if ("statusele = moving" || "statusele = stopped"){
        door = close;
    }
    else{
        opendoor();
    }
}
function findfloorbutton(){
    for (let floorbutton = 0; floorbutton < array.length; floorbutton++) {
    if ("currentfloor = floorbutton")
    return button();
    }
}
function checkelevatorstatus(){
    array.forEach(elevator =>column) ,{  
        if (dooropentimer = 0){
        closedoor();
        }
        if(idletimer = 0){
        addestination(defaultfloor,elevator,usermove)
    };
    }
}

function startmove(elevator){
    if ("currentfloor < destinationlist"){
        direction += 1;
    }else{
        direction -= 1;
    }
    statusele = moving
}

function stopelevator(elevator){
    setTimeout(() => {
        
    }, timeout);
    statusele = stop
}

function finddirectionbutton(direction, requestedfloor){
    array.forEach(button => directionbutton {
       if (requestedfloor = floorbutton && direction = directionbutton){
       return button;
       }
    });
}

function callelevator(directionbutton, requestedfloor, usermove){
    if (findelevator(direction, requestedfloor)){
     elevator()
    }
    else if (checkifdestinexist ((requestedfloor, elevator) = false)){
        addestination(elevator, requestedfloor, usermove)
    }
}

function checkmovingelevator(){
    array.forEach(elevator => column{
        if (statusele =moving){
            if (elevator, currentfloor < requestedfloor || currentfloor > requestedfloor){
            currentfloor /= floor;
            }else if (elevator, currentfloor = destinationlist){
                stopelevator(elevator) = true;
                clearbuttons(elevator);
            }
            opendoor(elevator);
        }
    });
}

function clearbuttons (elevator)
    finddirectionbutton(elevator, direction, currentfloor)
