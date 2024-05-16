package main

import "fmt"

func lastElement[T any](values []T) T {
	if len(values) == 0 {
		var t T
		return t
	}

	return values[len(values)-1]
}
func main() {

	s := []int{1, 2, 3, 4}
	r := []string{"1", "2", "3"}

	b := []bool{true, false, true, false}

	fmt.Println(lastElement(s), lastElement(r), lastElement(b))
}
