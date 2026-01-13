package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	// Check for division by zero
	if b == 0 {
		return 0, errors.New("Cannot Divide by Zero")
	}
	return a / b, nil
}
func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
