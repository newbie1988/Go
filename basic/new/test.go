package main

import (
	"fmt"
	"./common"
	)

func main() {
	p := common.NewPerson("Jack", "HangZhou", "Male", 21);
	// compile error.
	//p.name = "Mike";
	fmt.Println(p);

	//compile error.
	//p1 := common.person{name:"Lily", age:12, address:"ShangHai", gender:"Female"};
	//fmt.Println(p1);
}
