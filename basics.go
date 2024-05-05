package main

import "fmt"

func basics() {
	fmt.Println("Hello world")

	// vairable declaration
	var num int = 15

	// print
	fmt.Println(num)

	// shorthand variable declaration - except for constant
	penniesPerText := 2

	// print formated string
	fmt.Printf("Pennies per text is %T\n", penniesPerText)

	// multiple variables declaring
	myname, mynum := "ayush", 123

	// %T for type
	fmt.Printf("%T %T\n", myname, mynum)

	// typecasting
	myFloat := 3.99
	myInt := int(myFloat)
	myFloat = float64(myInt)
	fmt.Println(myFloat)

	//constant -> cannot use shorthand
	const myconst = 5 // cannot ise shorthand to declare it
	fmt.Println(myconst)

}
