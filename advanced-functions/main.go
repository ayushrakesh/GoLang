package main

import (
	"fmt"
	// "os"
)

type user struct {
	name  string
	phone int
	admin bool
}

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
	e := arithmetic(d, c)

	return e
}
func deleteUser(userMap map[string]user, name string) string {
	user, ok := userMap[name]

	if !ok {
		return "User not found!"
	}
	if user.admin {
		return "Admin deleted!"
	}
	defer delete(userMap, name)
	return "User deleted!"

}
func main() {
	addResult := aggregate(1, 2, 5, sum)
	mulResult := aggregate(1, 2, 10, mul)

	fmt.Println(addResult, mulResult)

	userMap := make(map[string]user)
	userMap = map[string]user{
		"john": {
			name: "john", phone: 123, admin: true,
		},
		"steve": {
			name: "steve", phone: 124, admin: false,
		},
		"mitchell": {
			name: "mitchell", phone: 196, admin: true,
		},
	}

	fmt.Println(deleteUser(userMap, "johny"))
	// deleteUser(userMap, "john")
	// fmt.Println(("john"))

	subResult := aggregate(10, 5, 2, func(i1, i2 int) int {
		return i1 - i2
	})
	fmt.Println(subResult)
}
