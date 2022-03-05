package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
	fmt.Println(greeting("John"))
}

func soma(a int, b int) int {
	return a + b
}

func greeting(name string) string {
	greeting := "Hello " + name
	return greeting
}
