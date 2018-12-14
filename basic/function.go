package main
import "fmt"

func max(a int, b int) int {
	var tmp int = 0;
	if a > b {
		tmp = a;
	} else {
		tmp = b;
	}
	return tmp;
}

func main() {
	a, b := 5, 9;
	fmt.Println("The max num is ", max(a, b));
}
