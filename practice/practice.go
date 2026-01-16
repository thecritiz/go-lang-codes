package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	var res int = add(8, 4)
	fmt.Println(res)
}
