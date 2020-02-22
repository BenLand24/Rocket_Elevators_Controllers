
class Battery{
    constructor(nbfloors,nbele){
        this.nbfloors = nbfloors;

        this.nbele = nbele;

        this.Column = new Column(nbfloors,nbele);
    }

    requestele(requestedfloornumber, direction){
        if (requestedfloornumber > this.nbfloors || requestedfloornumber < 0){
            return console.log(
                "this floor is not vailde, try again if it fails again, maybe you're in the wrong building?"
            );     
        } else {
            console.log("elevator requested at floor number", requestedfloornumber);
            sleep(2000);
            console.log(
                "elevator",
                this.Column.elelist[0].elenumber,
                "is currently at floor number",
                this.Column.elelist[0].currentfloor
            );
            sleep(2000);
            console.log(
                "elevator",
                this.Column.elelist[1].elenumber,
                "is currently at floor number",
                this.Column.elelist[1].currentfloor
            );
            sleep(2000);

            var nearele = this.findele(requestedfloornumber, direction);

            console.log("returning elevator", nearele.elenumber);
            sleep(2000);

            nearele.addfloorlist(requestedfloornumber);

            nearele.movenext();

            return nearele;
        }
    }
    requestfloor(elevator, requestedfloornumber){
        if (elevator === undefined){
            return;
        } else {
            elevator.activateinsidebutton(requestedfloornumber);

            elevator.addfloorlist(requestedfloornumber);

            elevator.movenext();
        }
    }
    findele(requestedfloornumber) {
        let bestdifference = 10;
        let nearele = null;

        for (let i = 0; i < this.Column.elelist.length; i++) {
            const differencefloor = Math.abs(
                requestedfloornumber - this.Column.elelist[i].currentfloor
            );
            if (differencefloor < bestdifference) {
                bestdifference = differencefloor;
                nearele = i;
            }
        }
        return this.Column.elelist[nearele];
    }
}
// Column
class Column{
    constructor(nbfloors, nbele){
        this.nbfloors = nbfloors;

        this.nbele = nbele;

        this.elelist = [];

        this.createeleator();
    }
    
    createeleator(){
        for (let i = 0; i < this.nbele; i++) {
            let elevator = new Elevator(i + 1, this.nbfloors);
            
            this.elelist.push(elevator);
        }
    }
}

//elevator
class Elevator{
    constructor(elenumber, nbfloors){
        this.elenumber = elenumber;

        this.nbfloors = nbfloors;

        this.direction = "stop";

        this.status = "idle";

        this.requestfloorlist = [];

        this.currentfloor = 0;
    }

    movenext(){
        let requestedfloornumber = this.requestfloorlist.shift();

        if(this.currentfloor === requestedfloornumber) {
            console.log(
                "elevator",
                this.elenumber,
                "arrived at floor number",
                this.currentfloor
            );
            sleep(2000),
            this.opendoor();
            this.closedoor();
        } else {
            while(this.currentfloor != requestedfloornumber) {
                if (this.currentfloor > requestedfloornumber){
                    this.movedown();
                    console.log("at", this.currentfloor, "move down");
                    sleep(2000);
                }else if (this.currentfloor < requestedfloornumber) {
                    this.moveup();
                    console.log("at", this.currentfloor, "move up");
                    sleep(2000);
                }
            }
            console.log(
                "elevator",
                this.elenumber,
                "arrived at floor number",
                this.currentfloor
            );
            sleep(2000);
            this.opendoor();

            this.closedoor();
        }
    }

    movedown(){
        this.direction = "down";

        this.status = "moving";

        this.currentfloor = this.currentfloor -1;
        sleep(2000);
    }

    moveup() {
        this.direction = "up";

        this.status = "moving";

        this.currentfloor = this.currentfloor +1;
        sleep(2000);
    }

    addfloorlist(requestedfloornumber) {
        this.requestfloorlist.push(requestedfloornumber);

        if (this.direction == "up") {
            this.requestfloorlist.sort();
        } else if (this.direction == "down") {
            this.requestfloorlist.sort().reverse();
        }
    }

    activateinsidebutton(requestedfloornumber) {
        console.log(
            "in elevator",
            this.elenumber,
            ", which is at floor",
            this.currentfloor,
            ",",
            "floor number",
            requestedfloornumber,
            "is requested"
        );
        sleep(2000);
    }
    opendoor() {
        console.log("opening doors at floor number", this.currentfloor);
        sleep(2000);
    }

    closedoor() {
        console.log("closing doors");
        sleep(2000)
    }
}

//timer

function sleep(milliseconds) {
    var start = new Date().getTime();
    for (let i = 0; i < 1e7; i++) {
        if (new Date().getTime() - start >milliseconds) {
            break;
        }        
    }
}

//Battery value
 Battery = new Battery(10, 2);

// test 1

Battery.Column.elelist[0].currentfloor = 2;

Battery.Column.elelist[1].currentfloor = 6;

requestelenumber1 = Battery.requestele(7, "down");

Battery.requestfloor(requestelenumber1, 5);

// test 2

// Battery.Column.elelist[0].currentfloor = 5;

// Battery.Column.elelist[1].currentfloor = 3;

// requestelenumber1 = Battery.requestele(1, "up");

// Battery.requestfloor(requestelenumber1,7);

// requestelenumber2 = Battery.requestele(3, "up");

// Battery.requestfloor(requestelenumber2, 5);

// requestelenumber1 = Battery.requestele(9, "down");

// Battery.requestfloor(requestelenumber1, 1);