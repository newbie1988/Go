package main
import (
	"fmt"
	"time"
)

func calc(ch chan int, num int) {
	go func() {
		time.Sleep(time.Duration(num) * time.Second);
		ch <- num;
	}();
}

func main() {
	ch := make(chan int);
	calc(ch, 1);
	calc(ch, 2);
	num1 := <- ch;
	num2 := <- ch;
	fmt.Printf("result is %d\n", num1 * num2);
}
