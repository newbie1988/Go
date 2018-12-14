package main
import (
	 "fmt"
	 "strings"
	)

func addSuffix(suffix string) func(string) string {
	return func(name string) string {
			if !strings.HasSuffix(name, suffix) {
				name += suffix;
			}
			return name;			
		};
}

func main() {
	bmpSuffix := addSuffix(".bmp");
	fmt.Println("bmp: ", bmpSuffix("aaa"));
	fmt.Println("bmp: ", bmpSuffix("ddd.bmp"))
	jpgSuffix := addSuffix(".jpg");
	fmt.Println("jpg: ", jpgSuffix("bbbb"));
	fmt.Println("jpg: ", jpgSuffix("cccc.jpg"));
}
