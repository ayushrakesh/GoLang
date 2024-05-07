package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func understandingSlices() {
	myArray := [5]int{1, 2}
	mySlices := myArray[2:]
	fmt.Println(mySlices)
	fmt.Printf("%T\n", mySlices)

	primes := []int{2, 3, 5, 7}
	fmt.Println(primes)
	fmt.Printf("%T\n", primes)

	s := []string{"I", "love", "you"}
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	t := make([]bool, 4)
	fmt.Println(t)
	fmt.Printf("%T\n", t)
}
func appendSlices(s []int) []int {
	t := []int{6, 7}
	s = append(s, t...)
	return s
}

func bucket(costs []cost) []float64 {
	myBucket := []float64{}
	fmt.Println("----------------------")
	fmt.Println(myBucket, len(myBucket))

	for i := 0; i < len(costs); i++ {
		costOfDay := 0.0
		for j := 0; j < len(costs); j++ {
			if costs[j].day == i {
				costOfDay += costs[j].value
			}
		}

		if i != len(myBucket) {
			myBucket = append(myBucket, 0.0)
			continue
		} else {
			myBucket = append(myBucket, costOfDay)
		}

	}
	return myBucket
}
func createMatrix() {
	rows := 10
	columns := 10
	myMatric := [][]int{}
	for i := 0; i < rows; i++ {
		row := make([]int, columns)

		for j := 0; j < columns; j++ {
			row = append(row, i*j)
		}
		myMatric = append(myMatric, row)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%d ", myMatric[i][j])
		}
		fmt.Printf("\n")
	}

}
func understandingCapacity(s []int) {

	s[0] = 10

}
func main() {
	understandingSlices()

	s := []int{1, 2, 3}
	s = appendSlices(s)

	fmt.Println(s)

	c := []cost{
		{
			day: 0, value: 1.2,
		},
		{
			day: 1, value: 5.0,
		},
		{
			day: 5, value: 6.2,
		},
		{
			day: 7, value: 3.2,
		},
	}

	fmt.Println(bucket(c))

	createMatrix()

	t := make([]int, 3, 8)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	// fmt.Println(t)

	w := t[1:]
	w[0] = 1

	x := append(w, 2, 3, 4, 5, 6, 7)
	x[0] = 2
	fmt.Println(x)
	fmt.Printf("%v %v %v\n", &t[0], &w[0], &x[0])
	understandingCapacity(t)

	fmt.Println(len(t), cap(t))
	fmt.Println(len(w), cap(w))
	fmt.Println(len(x), cap(x))
}
