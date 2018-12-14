package main
import "fmt"

func printLevel(score int) {
	switch score {
		case 90:
			fmt.Println("A");
		case 80:
			fmt.Println("B");
		case 70:
			fmt.Println("C");
		default:
			fmt.Println("D");
		}
}

func main() {
	printLevel(90);
	printLevel(70);
	printLevel(50);
}
