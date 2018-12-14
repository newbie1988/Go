package main
import "fmt"

type Book struct {
	name string;
	id int32;
	price float64;
	author string;
}

func (book Book) getName() string {
	fmt.Println("Book Name: " + book.name);
	return book.name;
}

func (book Book) printDetails() {
	fmt.Println(book);
}

func (book *Book) setPrice(price float64) {
	book.price = price;
}

func main() {
	book := Book{"Pattern Design", 11111, 30.5, "GoF"};
	book.getName();
	book.printDetails();
	(&book).setPrice(121);
	book.printDetails();
	
	var newBook Book = Book{name : "Thinking in Java", author : "Bruce Eckel", price : 70};
	newBook.printDetails();
}
