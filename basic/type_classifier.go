package main
import "fmt"

type Shape interface {
	Whoami();
}

type Circle struct {
}

func (circle Circle) Whoami() {
	fmt.Println("I am Circle");
}

type Rectangle struct {
}

func (r Rectangle) Whoami() {
	fmt.Println("I am Rectangle");
}

type Triangle struct {
}

func (triangle Triangle) Whoami() {
	fmt.Println("I am Triangle");
}
func ShowTypeInfo(shape Shape) {
	switch shape.(type) {
		case Circle, *Circle :
				fmt.Println("Type: Circle");
		case Rectangle, *Rectangle:
				fmt.Println("Type: Rectangle");
		case nil:
				fmt.Println("It's nil");
		default:
				fmt.Println("Unkonwn Type");
	}
}

func ParserTypeInfo(types...interface{}) {
	for _, t := range types {
		switch t.(type) {
			case int, int64:
				fmt.Println(t, "Integer");
			case bool:
				fmt.Println(t, "Boolean");
			case string:
				fmt.Println(t, "String");
			case float32, float64:
				fmt.Println(t, "Float");
			case nil:
				fmt.Println(t, "nil");
			default:
				fmt.Println(t, "Unkown Type");
		}
	}
}

func main() {
	var shape Shape = Circle{};
	ShowTypeInfo(shape); // Type: Circle
	shape.Whoami();
	shape = new(Rectangle);
	ShowTypeInfo(shape);
	shape.Whoami();
	shape = nil;
	ShowTypeInfo(shape); // It's nil
	var circle Circle;
	ShowTypeInfo(circle); // Type: Circle
	var triangle Triangle;
	ShowTypeInfo(triangle); // "Unkonwn Type
	fmt.Println("----------------------------");	
	ParserTypeInfo(1, 1.567, true, nil, "hello");
}
