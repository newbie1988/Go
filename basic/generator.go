package main
import "fmt"

type Object interface {};
type Maker func(Object) (Object, Object);

func GeneratorFactory(maker Maker, initial Object) func() Object {
	ch := make(chan Object);
	go func() {
		oldValue := initial;
		var newValue Object;
		for {
			newValue, oldValue = maker(oldValue);
			ch <- newValue;	
		}
	}();
	
	return func() Object {
		return <- ch;
	};
} 

func main() {
	even := func(num Object) (Object, Object) {
		oldValue := num.(int);
		newValue := oldValue + 2;
		return oldValue, newValue;
	}
	
	evenGenerator := GeneratorFactory(even, 0);
	for i := 0 ; i < 10; i++ {
		fmt.Println(evenGenerator());
	}
}
