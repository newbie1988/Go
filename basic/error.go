package main
import (
	"fmt"
	"errors"
)

func div(a, b int) (int, error) {
	if 0 == b {
		return 0, errors.New("div: dividend is zero");
	}
	
	return  a / b, nil;
}

func mod(a, b int) (int, error) {
	if(0 == b) {
		return 0, fmt.Errorf("div: mod is %d\n", b);
	}
	return a % b, nil;
}
func main() {
	if a, err := div(2, 0); err != nil {
		fmt.Println(err.Error());
	} else {
		fmt.Println(a);
	}
	
	if a, err := div(9, 3); err != nil {
		fmt.Println(err.Error());
	} else {
		fmt.Println(a);
	}

	if a, err := mod(8, 0); nil != err {
		fmt.Println(err.Error());
	} else {
		fmt.Println(a);
	}
}
