package main
import "fmt"

func main() {
	first := []int{1, 2, 3, 4, 5};
	second := []int{6, 7, 8, 9, 10};
	final := append(first, second...);
	for _, value := range final {
		fmt.Println(value);
	}
	
	firstBak := make([]int, 10);
	copyLen := copy(firstBak, first);
	fmt.Println("copy element nums:", copyLen);
	for _, value := range firstBak {
		fmt.Println(value);
	}
}
