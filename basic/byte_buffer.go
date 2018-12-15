package main
import (
	"fmt"
	"bytes"
	)

func main() {
	var buffer bytes.Buffer;
	buffer.WriteString("hello, ");
	buffer.WriteString("how ");
	buffer.WriteString("are ");
	buffer.WriteString("you ?");
	fmt.Println(buffer.String());
}
