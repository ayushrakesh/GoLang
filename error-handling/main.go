package main

import (
	"errors"
	"fmt"
)

func sendMsg(msg string) (string, error) {
	if len(msg) <= 8 {
		return msg, nil
	} else {
		return "Error", fmt.Errorf("Something went wrong!")
	}
}
func divide(dividend, divisor float32) (float32, error) {
	if divisor == 0 {

		// standard library package function -> errors.New()
		return 0.0, errors.New("no divide by zero")
	}
	return dividend / divisor, nil
}
func calculateMaxMessages() float64 {
	thre := 100

	// maxMessages:=0
	totalCost := 0.0

	for i := 0; ; i++ {
		totalCost = totalCost + (1.0 + float64(i)*0.01)

		if totalCost > float64(thre) {
			return float64(i)
		}
	}
}
func understandingLoops() {

	maxMessages := calculateMaxMessages()
	fmt.Println(maxMessages)
}
func main() {
	msg, err := sendMsg("this is")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	// fmt.Println(msg)

	fmt.Println(divide(5, 7))

	understandingLoops()
}
