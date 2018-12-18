package common 

type person struct {
	name string;
	age int;
	address string;
	gender string;
}

func NewPerson(name, address, gender string, age int) *person {
	return &person{name, age, address, gender};
}
