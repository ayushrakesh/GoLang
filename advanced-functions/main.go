package main

import "fmt"

//first class function
func sum(a, b int) int {
	return a + b
}

// first class function
func mul(a, b int) int {
	return a * b
}

// higher-order function
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	d := arithmetic(a, b)
	e := arithmetic(c, d)

	return e
}
func main() {
	addResult := aggregate(1, 2, 5, sum)
	mulResult := aggregate(1, 2, 10, mul)

	fmt.Println(addResult, mulResult)

}
