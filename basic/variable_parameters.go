package main
import "fmt"

func min(nums...int) int {
	if 0 == len(nums) {
		return 0;
	}

	minNum := nums[0];
	for index := 1; index < len(nums); index++ {
		if minNum > nums[index] {
			minNum = nums[index];
		}
	}
	return minNum;
}

func main() {
	fmt.Printf("The minimum num is: %d\n", min(5, 89, 1, 222, 0, 2));
	nums := []int{8, 2, 888, 10, 15, 28, 77};
	fmt.Printf("The minimum is: %d\n", min(nums...));
}
