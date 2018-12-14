package main
import "fmt"

func f() (ret int) {
	defer func() {
		fmt.Printf("ret is %d\n", ret);
		ret++;
	}();
	return 1;
}

func main() {
	i := 2;
	func() {
		i += 2;
	}();
	fmt.Printf("i = %d\n", i);
	fmt.Println("result is ", f());
}
