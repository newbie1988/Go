package main
import "fmt"

func crashed() {
	panic("program crahed");
}

func test() {
	fmt.Println("Enter test()");
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err);
		}
	}();
	crashed();
	fmt.Println("Leave test()");
}

func call() {
	fmt.Println("Eneter call()");
	test();
	fmt.Println("Leave call()");
}

func main() {
	fmt.Println("Start Main");
	call();
	fmt.Println("End Main");
}
