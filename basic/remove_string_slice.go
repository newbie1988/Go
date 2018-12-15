package main
import "fmt"

func removeString(strs []string, start int, end int) []string{
	length := len(strs);
	if start < 0 || end > length {
		fmt.Println("invalid paramters!");
		return nil;
	}

	j := 0;
	for i := start; i < length && j < length - end; i++ {
		strs[i] = strs[end + j];
		j++;
	}
	
	strs = strs[0:length - end + start];
	return strs;
}

func main() {
	strs := []string{"111", "2222", "3333", "4444", "5555", "6666", "7777"};
	//strs = removeString(strs, 0, 2);
	//strs = removeString(strs, 2, len(strs));
	strs = removeString(strs, 3, 5);
	for _, value := range strs {
		fmt.Println(value);
	}
}
