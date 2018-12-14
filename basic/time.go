package main
import (
	"fmt"
	"time"
)

func main() {
	t := time.Now();
	fmt.Println("now is", t);
	fmt.Printf("%d_%d_%d\n", t.Year(), t.Month(), t.Day());
	fmt.Println("UTC time: ", t.UTC());
	//note "2006 02 Jan 15:04" the content is fixed, should never change content.
	fmt.Println("Format time: ", t.Format("2006 02 Jan 15:04"));
	fmt.Println("Format time: ", t.Format("20060102"));
	var week time.Duration = 7 * 24 * 60 * 60 * 1e9;
	now := time.Now();
	fmt.Println("now is: ", now);
	weekFromNow := now.Add(week);
	fmt.Println("after a week, time is: ", weekFromNow);
}
