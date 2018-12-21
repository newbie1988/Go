package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	name string;
	brand string;
	id int;
}

func (car Car) String() string {
	return car.name + ":" + car.brand + ":" + strconv.Itoa(car.id); 
}

func (this *Car) ModifyName(newName string) {
	this.name = newName;
}

func main() {
	car := Car{"BMW", "BMW-X7", 111};
	fmt.Println(car);
	car.ModifyName("Audio");
	fmt.Println(car);
	
	newCar := &Car{"Benz", "C200", 6666};
	fmt.Println(newCar);
	newCar.ModifyName("BYD");
	fmt.Println(newCar);
}
