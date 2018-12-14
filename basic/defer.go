package main
import "fmt"

func trace(s string) string {
	fmt.Printf("Enter %s\n", s);
	return s;
}

func untrace(s string) {
	fmt.Printf("Leave %s\n", s);
}

func first() {
	defer untrace(trace("first"));
	fmt.Println("In Function first");
}

func second() {
	defer untrace(trace("second"));
	fmt.Println("In Function second");
	first();
}

func main() {
	second();
}
