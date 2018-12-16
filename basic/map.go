package main
import "fmt"

func print(weathers map[string]string) {
	for city, weather := range weathers {
		fmt.Printf("%s------->%s\n", city, weather);
	}
}
func main() {
	weathers := map[string]string{"HangZhou":"snow", "BeiJin":"rain", "ShangHai":"clound"};
	print(weathers);
	weathers["HangZhou"] = "sunny";
	fmt.Println("++++++++++++++++++");
	print(weathers);
	fmt.Println("##################");
	delete(weathers, "HangZhou");
	if weather, isExist := weathers["HangZhou"]; isExist {
		fmt.Println("Hangzhou weather is", weather);
	} else {
		fmt.Println("Can't find HangZhou Weather");
	}
	print(weathers);
}
