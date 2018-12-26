package main
import (
	"fmt"
	"flag"
)

var bNewLine = flag.Bool("n", false, "add new line");
var num = flag.Int("d", 0, "nums:");

func main() {
	flag.PrintDefaults();
	flag.Parse();
	var str string;
	for i := 0; i < flag.NArg(); i++ {
		str += flag.Arg(i);
		if *bNewLine {
			str += "\n"
		}
	}
	if nil != num {
		fmt.Println(*num);
	}
	fmt.Println(str);
}
