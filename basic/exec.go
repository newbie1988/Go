package main
import (
//	"fmt"
	"os"
	"os/exec"
//	"bytes"
)

func main() {
	cmd := exec.Command("ls", "-l");
	//var out bytes.Buffer;
	cmd.Stdout = os.Stdout;
	cmd.Run();
//	fmt.Println(out.String());
}
