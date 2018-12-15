package main
import "fmt"

func print(slice []int) {
	fmt.Println("--------------------");
	for _, value := range slice {
		fmt.Println(value);
	}
}

func main() {
	slice := make([]int, 5, 10);
	for index := range slice {
		slice[index] = index;
	}
	print(slice);
	
	length := len(slice);
	slice = slice[0:length + 2];
	slice[length] = 88;
	slice[length + 1] = 100;
	print(slice);
	slice = slice[0:0];
	fmt.Println("length of slice[0:0] is", len(slice));	
}
