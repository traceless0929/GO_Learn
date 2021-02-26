package main

import "fmt"

func init() {
	fmt.Printf("Hello Init")
}

func main() {
	fmt.Printf("\n%d", add(1, 2))
}

func add(a int, b int) int {
	return a + b
}
