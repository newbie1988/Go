package main
import "fmt"

type Sum int;

func (sum Sum) Add(num Sum) Sum {
	return sum + num;
}

func main() {
	var a Sum = 5;
	var b Sum = 6;
	fmt.Println("result is", a.Add(b));
}
