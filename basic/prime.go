package main
import "fmt"

func generator(ch chan int) {
	for i := 2; ; i++ {
		ch <- i;
	}
}

func filter(in, out chan int, prime int) {
	for {
		num := <- in;
		if num % prime != 0 {
			out <- num;
		}
	}
}

func main() {
	ch := make(chan int);
	go generator(ch);
	for {
		prime := <- ch;
		fmt.Println(prime);
		ch1 := make(chan int);
		go filter(ch, ch1, prime);
		ch = ch1;
	}
}
