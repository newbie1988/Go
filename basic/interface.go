package main
import "fmt"

type Car interface {
	run();
}

type BMW struct {
	name string;
}

type Benz struct {
	name string;
}

func (car BMW) run() {
	fmt.Printf("%s is running\n", car.name);
}

func (car Benz) run() {
	fmt.Printf("%s is running\n", car.name);
}

func main() {
	var car Car = nil;

	bmw := BMW{name : "BMW"};
	benz := Benz{name : "Benz"};
	car = &bmw;
	car.run();
	car = &benz;
	car.run();

	car = &BMW{"BMW"};
	car.run();
}
