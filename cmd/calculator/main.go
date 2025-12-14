// Package main demonstrates the usage of the calculator package.
package main

import (
	"errors"
	"fmt"

	"github.com/michael-freling/claude-code-sandbox/internal/calculator"
)

func main() {
	fmt.Println("Calculator Demo")
	fmt.Println("===============")
	fmt.Println()

	a, b := 10, 5

	sum := calculator.Add(a, b)
	fmt.Printf("Add(%d, %d) = %d\n", a, b, sum)

	diff := calculator.Subtract(a, b)
	fmt.Printf("Subtract(%d, %d) = %d\n", a, b, diff)

	product := calculator.Multiply(a, b)
	fmt.Printf("Multiply(%d, %d) = %d\n", a, b, product)

	quotient, err := calculator.Divide(a, b)
	if err != nil {
		fmt.Printf("Divide(%d, %d) failed: %v\n", a, b, err)
	} else {
		fmt.Printf("Divide(%d, %d) = %d\n", a, b, quotient)
	}

	fmt.Println()
	fmt.Println("Error Handling Demo")
	fmt.Println("===================")
	fmt.Println()

	_, err = calculator.Divide(a, 0)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			fmt.Printf("Divide(%d, 0) correctly returned division by zero error: %v\n", a, err)
		} else {
			fmt.Printf("Divide(%d, 0) returned unexpected error: %v\n", a, err)
		}
	}
}
