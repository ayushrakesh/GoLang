package main

import "fmt"

func functions() {
	func concat(firstName, lastName string) string {
		return firstName + " " + lastName
	}
	func updateX(x int) int { // should have to return value to update variables
		x = x + 1
		return x
	}
	func nakedReturn(mystring string) (s1, s2, s3 string) {
		s1 += mystring
		s2 += mystring
		s3 += mystring
		return
	}
}
