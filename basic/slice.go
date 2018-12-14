package main
import "fmt"

func print(elems []int) {
	fmt.Println("length is ", len(elems));
	fmt.Println("capacity is ", cap(elems));
	for _, elem := range elems {
		fmt.Println(elem);
	}
}

func main() {
	elems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9};
	print(elems);
	fmt.Println("++++++++++++++++++++");
	print(elems[:5]);
	fmt.Println("++++++++++++++++++++");
	print(elems[3:]);
	fmt.Println("++++++++++++++++++++");
	print(elems[5:8]);
}
