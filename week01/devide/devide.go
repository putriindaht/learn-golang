package divide

import "fmt"

func Bagi(numerator, denominator int) int {
	if denominator == 0 {
		fmt.Println("Division by zero is not allowed")
		return 0
	}

	if numerator < denominator {
		fmt.Printf("Cannot devide: %d is less then %d (numerator < denominator)\n", numerator, denominator)
		return 0
	}

	return numerator / denominator
}
