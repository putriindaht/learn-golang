package main

import (
	"fmt"

	"github.com/putriindah/learn-golang-week01/add"
	divide "github.com/putriindah/learn-golang-week01/devide"
	"github.com/putriindah/learn-golang-week01/variadicSubtract"
)

func main() {
	// Add Number
	addedNumber := add.Add(1, 2)
	fmt.Println(addedNumber, "-> added")

	// Divide Number
	devidedNumber := divide.Bagi(2, 6)
	fmt.Println(devidedNumber, "-> divided")

	// Variadic
	subtracted := variadicSubtract.KurangVariadic(20, 2, 1, 2, 1)
	fmt.Println(subtracted, "-> substracted")

}
