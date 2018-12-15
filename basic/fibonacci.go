package main

import (
	"fmt"
	"time"
	)

func fibonacci(n int) int {
	if n <= 1 {
		return 1;
	}
	
	return fibonacci(n - 1) + fibonacci(n - 2);
}

const LIM int = 30;
var fibocis [LIM]int64;

func fibonacciCache(n int) int64 {
	if n < LIM && fibocis[n] != 0 {
		return fibocis[n];
	}
	
	if n <= 1 {
		return 1;
	}
	return fibonacciCache(n - 2) + fibonacciCache(n - 1);
}

func main() {
	start := time.Now();
	for i := 0; i < LIM; i++ {
		fibonacci(i);
	}
	diff := time.Now().Sub(start);
	fmt.Printf("fibonacci witch no cache cost time: %s\n", diff);
	
	start = time.Now();
	for i := 0; i < LIM; i++ {
		fibocis[i] = fibonacciCache(i);
	}
	diff = time.Now().Sub(start);
	fmt.Printf("fibonacci with cache cost time: %s\n", diff);
}
