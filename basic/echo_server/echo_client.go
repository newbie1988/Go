package main
import (
	"fmt"
	"net"
	"flag"
	"bufio"
	"os"
	"strings"
)

func main() {
	flag.Parse();
	if 2 != flag.NArg() {
		panic("Usage: <IP> <Port>");
	}
	serverIP := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1));
	con, err := net.Dial("tcp", serverIP);
	if nil != err {
		panic("connnect failed");
	}
	defer con.Close();
	input := bufio.NewReader(os.Stdin);
	for {
		data, _ := input.ReadString('\n');
		data = strings.Trim(data, "\n");
		if "Q" == data {
			fmt.Println("Quit");
			break;
		}
		_, err := con.Write([]byte(data));
		if nil != err {
			fmt.Println("write error:", err.Error());
			break;
		}
	}
}
