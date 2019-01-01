package main
import "fmt"

const COUNT = 10;

func main() {
	counts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	results := make([]int, COUNT);
	ch := make(chan int);	

	for index, count := range counts {
		go func(count int, index int) {
			for loop := 0; loop < count; loop++ {
				for i := 1; i <= 100; i++ {
					results[index] += i;
				} 
			}
			ch <- 0;
		}(count, index);
	}
	
	for i := 0; i < COUNT; i++ {
		<- ch;
	}
	
	for _, value := range results {
		fmt.Println(value);
	}
}
