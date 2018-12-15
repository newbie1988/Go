package main
import (
	"fmt"
	"time"
	)

func calculate() {
	sum := 0;
	for i := 0; i < 100000; i++ {
		sum += i;
	}
	//fmt.Printf("sum is %d\n", sum);
}

func main() {
	start := time.Now();
	for i := 0; i < 1000000; i++ {
		calculate();
	}
	diff := time.Now().Sub(start);
	fmt.Printf("cost time: %s\n", diff);
}
