package main
import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin);
	for {
		line, err := reader.ReadString('\n');
		if(nil == err) {
			line = strings.Trim(line, "\n");
			line = strings.Replace(line, "吗", "", -1);
			line = strings.Replace(line, "？", "!", -1);
			line = strings.Replace(line, "?", "!", -1);
			fmt.Println(line);
		}
	}
}
