package main
import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6};
	for _, num := range nums {
		fmt.Println(num);
	}
	
	for index, num := range nums {
		fmt.Printf("%d----->%d\n", index, num);
	}
}
