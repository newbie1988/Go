package main
import (
	"fmt"
	"time"
)

func produce(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i;
		fmt.Println(i);
	}
}

func main() {
	ch := make(chan int);
	go produce(ch);
//	fmt.Println(<-ch);
	time.Sleep(5 * 1e9);
}
