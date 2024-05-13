package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {

	go func(t int) {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received- %v\n", message)
	}(250)
	fmt.Printf("Email sent- %v\n", message)

}
func main() {
	sendEmail("Email 1")
	sendEmail("Email 2")
	sendEmail("Email 3")
	sendEmail("Email 4")
}
