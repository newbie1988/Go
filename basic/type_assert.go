package main
import (
	"fmt"
)

type Car interface {
	Whoami();
}

type BMW struct {
}

type Benz struct {
}

type Audio struct {
}
func (bmw BMW) Whoami() {
	fmt.Println("I am BMW");
}

func (benz Benz) Whoami() {
	fmt.Println("I am Benz");
}

func (audio Audio) Whoami() {
	fmt.Println("I am Audio");
}

func PrintMsg(car Car) {
	if _, ok := car.(BMW); ok {
		fmt.Println("Type is BMW");		
	} else if _, ok := car.(Benz); ok {
		fmt.Println("Type is Benz");
	} else {
		fmt.Println("Unkown Type!!");
	}
}
func main() {
	var car Car = BMW{};
	PrintMsg(car);
	car.Whoami();
	car = Benz{};
	PrintMsg(car);	
	car.Whoami();
	car = Audio{};
	PrintMsg(car);
	car.Whoami();
}
