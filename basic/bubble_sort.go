package main
import "fmt"

func bubbleSort(elems []int, size int) {
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < (size - 1 - i); j++ {
			if elems[j] > elems[j + 1] {
				elems[j], elems[j + 1] = elems[j + 1], elems[j];
			}
		}
	}
}

func main() {
	nums := []int{9, 4, 5, 8, 10, 2, 6, 9};
	fmt.Println("Before Bubble Sort: ", nums);
	bubbleSort(nums, len(nums));
	fmt.Println("After Bubble Sort: ", nums);
}
