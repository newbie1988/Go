package main
import (
	"fmt"
	"time"
)

func producor1() chan int {
	ch := make(chan int);
	go func() {
		for i := 0; ; i++ {
			ch <- 2 * i;
		}
	}();
	return ch;
}

func producor2() chan int {
	ch := make(chan int);
	go func() {
		for i := 0; ; i++ {
			ch <- (9 * i + 1);
		}
	}();
	return ch;
}

func consumer() {
	ch1 := producor1();
	ch2 := producor2();
	go func() {
		for {
			select {
				case v := <- ch1:
					fmt.Println("channel 1, value:", v);
				case v := <- ch2:
					fmt.Println("channel 2, value:", v);
/*				default:
					fmt.Println("no channel ready!");
*/
			}
		}
	}();
}

func main() {
	consumer();
	time.Sleep(1e9);
}
