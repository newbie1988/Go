package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

const (
	InputFileName = "test_in.data"
	OutputFileName = "test_out.data"
)

func main() {
	input, inErr := os.Open(InputFileName);
	if(nil != inErr) {
		fmt.Println("Open", InputFileName, "Failed!");
		return;
	}
	defer input.Close();
	
	inputReader := bufio.NewReader(input); 
	output, outErr := os.OpenFile(OutputFileName, os.O_WRONLY | os.O_CREATE, 0666);
	if(nil != outErr) {
		fmt.Println("Open", OutputFileName, "Failed!");
		return;
	}
	defer output.Close();
	outputWriter := bufio.NewWriter(output);
	
	for {
		line, err := inputReader.ReadString('\n');
		if(err == io.EOF) {
			break;
		}	
		outputWriter.WriteString(line);
	}
	// need flush before close. flush the buffer to file.
	outputWriter.Flush();
}
