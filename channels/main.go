package main

import (
	"fmt"
	"time"
)

type email struct {
	t time.Time
}

func checkTime(emails []email) {
	ch := make(chan bool)

	go func() {
		for _, em := range emails {
			if em.t.Before(time.Date(2024, 1, 2, 12, 0, 0, 0, time.Local)) {
				ch <- true
				continue
			} else {
				ch <- false
			}
		}
	}()

	// for i := 0; i < len(emails); i++ {
	v := <-ch
	fmt.Println(v)

	v = <-ch
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
}

func bufferChannelsWriting(emails []string) chan string {
	ch := make(chan string, len(emails))

	for i := 0; i < len(emails); i++ {
		ch <- emails[i]
		fmt.Printf(emails[i])
	}

	return ch
}
func bufferChannelsReading(emails []string, ch chan string) {
	for i := 0; i < len(emails); i++ {
		// <-ch
		fmt.Printf("%v\n", <-ch)
	}
}
func concurrentFib(n int) []int {
	ch := make(chan int)
	s := []int{}
	go fibonacci(n, ch)

	for item := range ch {
		s = append(s, item)
	}
	// _, ok := <-ch
	// fmt.Println(ok)
	return s
}
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	//closing from sender side
	close(ch)

}
func main() {

	emails := []email{
		{
			t: time.Date(2024, 1, 1, 12, 0, 0, 0, time.Local),
		},
		{
			t: time.Date(2024, 1, 2, 12, 0, 0, 0, time.Local),
		},
		{
			t: time.Date(2024, 1, 2, 12, 0, 0, 0, time.Local),
		},
		{
			t: time.Date(2024, 1, 2, 12, 0, 0, 0, time.Local),
		},
	}
	mails := []string{
		"mail 1",
		"mail 2",
		"mail 3",
		"mail 4",
	}

	checkTime(emails)

	ch := make(chan string)

	ch = bufferChannelsWriting(mails)
	fmt.Println("----------")
	bufferChannelsReading(mails, ch)

	fmt.Println(concurrentFib(10))

	cch := make(chan int)
	close(cch)
	// cch <- 5
	<-cch
	// fmt.Println(<-cch)

}
