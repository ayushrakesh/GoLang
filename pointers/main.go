package main

import (
	"fmt"
	"strings"
)

type person struct {
	gender string
}

func updateMsg(msgPtr *string) {
	// *msgPtr = "updatedMsg"

	// msg := *msgPtr
	*msgPtr = strings.ReplaceAll(*msgPtr, "aa", "****")
	// *msgPtr = msg
}
func updateStruct(p *person) {
	p.gender = "female"
}
func main() {
	message := "aaa"
	msgPtr := &message
	// *msgPtr = "world"

	updateMsg(msgPtr)
	fmt.Println(message)

	p := person{
		gender: "male",
	}
	updateStruct(&p)
	fmt.Println(p)
}
