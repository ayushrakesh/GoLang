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
	// }

}
func main() {
	// e := email{
	// 	t: time.Date(2024, 1, 1, 12, 0, 0, 0, time.Local),
	// }

	// fmt.Println(e.t)

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

	checkTime(emails)
}
