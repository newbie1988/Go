package main
import "fmt"

func A() {
	defer C();
	B();
}

func B() {
	defer D();
	panic("program crashed!\n");
}

func C() {
	fmt.Println("C()");
}

func D() {
	fmt.Println("D()");
}

func main() {
	A();
}
