package main

import "fmt"

func sendMsg(msg string) (string, error) {
	if len(msg) <= 8 {
		return msg, nil
	} else {
		return "Error", fmt.Errorf("Something went wrong!")
	}
}
func main() {
	msg, err := sendMsg("this is")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
}
