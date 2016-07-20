/*
	an example program
	输入流: fmt.Scanf
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("enter a number: ")

	var input float64
	fmt.Scanf("%f", &input)

	out := input * 2

	fmt.Println(out)
}
