package main
import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var str string = "This is an test";
	prefix := "this";
	fmt.Printf("has prefix:%s ? %t\n", prefix, strings.HasPrefix(str, prefix));
	suffix := "test";
	fmt.Printf("has suffix:%s ? %t\n", suffix, strings.HasSuffix(str, suffix));
	substr := "is";
	fmt.Printf("\"%s\" is substring of \"%s\"? %t\n", substr, str, strings.Contains(str, substr))
	fmt.Printf("\"is\" at the postion: %d\n", strings.Index(str, "is"));
	fmt.Printf("\"is\" at the last postion: %d\n", strings.LastIndex(str, "is"));
	fmt.Printf("Lower: %s\n", strings.ToLower(str));
	fmt.Printf("Upper: %s\n", strings.ToUpper(str));
	spaceStr := " hello world ";
	fmt.Printf("\"%s\" after trim is \"%s\"\n", spaceStr, strings.TrimSpace(spaceStr));
	strs := strings.Split(str, " ");
	for _, elem := range strs {
		fmt.Println(elem);
	}
	var num int = 25;
	fmt.Printf("%d the string format is \"%s\"\n", num, strconv.Itoa(num));
	numStr := "99999";
	cnt, _ := strconv.Atoi(numStr);
	fmt.Printf("\"%s\" the integer format is %d\n", numStr, cnt);
}
