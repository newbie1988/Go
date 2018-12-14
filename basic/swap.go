package main
import "fmt"

func swap(a *int, b *int) {
	*a ^= *b;
	*b ^= *a;
	*a ^= *b;
}

func main() {
	a, b := 8, 9;
	fmt.Printf("Before swap, a = %d, b = %d\n", a, b);
	swap(&a, &b)
	fmt.Printf("After swap, a = %d, b = %d\n", a, b);
	a, b = b, a;
	fmt.Printf("Restore value, a = %d, b = %d\n", a, b);
}
