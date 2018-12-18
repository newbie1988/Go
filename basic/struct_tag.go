package main

import (
	"fmt"
	"reflect"
	)

type Node struct {
	//should not contain semicolon
	data int "store value"
	next *Node "next node address"
}

func main() {
	node := Node{12, nil};
	for i := 0; i < 2; i++ {
		tagType := reflect.TypeOf(node);
		field := tagType.Field(i);
		fmt.Println(field.Tag);
	}
}
