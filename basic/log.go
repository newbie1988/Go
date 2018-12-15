package main
import (
	"fmt"
	"log"
	"runtime"
       )

func main() {
	debug := func() {
				_, file, line, _ := runtime.Caller(1);
				log.Printf("%s:%d\n", file, line);
			}
	fmt.Println("hello");
	debug();
}
