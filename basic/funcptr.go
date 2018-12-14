package main
import "fmt"

func sum(a int, b int) int {
	return a + b;
}

func callBack(a int, b int, f func (int, int) int) {
	fmt.Println("result is ", f(a, b));
}

func main() {
	callBack(2, 8, sum);
}
