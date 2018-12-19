package main
import "fmt"

type Car struct {
	name string;
	brand string;
}

type BMW struct {
	Car;
	speed int;
}

type Audio struct {
	Car;
	name string;
} 

func main() {
	car := BMW{Car{"BMW", "X7"}, 200};
	fmt.Println(car);
	car.name = "Benz";
	car.brand = "C300";
	car.speed = 220;
	fmt.Println(car);

	audio := Audio{Car{"Audio", "A8"}, "Audio-A8"};
	fmt.Println(audio);
	audio.name = "Audio-A6";
	fmt.Println(audio);
}
