package main

import "fmt"

type car struct {
	model string
	make  string

	frontWheel wheel
	backWheel  wheel
}
type vehicle struct {
	car
	regNo  int
	wheels wheel
}
type wheel struct {
	radius   float32
	material string
}

func concatenate(firstName, lastName string) string {
	return firstName + " " + lastName
}
func updateX(x int) int {
	x = x + 1
	return x
}
func nakedReturn(mystring string) (s1, s2, s3 string) {
	s1 += mystring
	s2 += mystring
	s3 += mystring
	return
}
func functions() {

}
func structs() vehicle {
	car1 := car{
		make:  "hyundai",
		model: "model 1",

		frontWheel: wheel{
			radius:   2.0,
			material: "steel",
		},
		backWheel: wheel{
			radius: 3.0, material: "iron"},
	}

	vehicle1 := vehicle{
		car:   car1,
		regNo: 100,
		wheels: wheel{
			radius:   5,
			material: "fibre",
		},
	}

	return vehicle1
}
func main() {

	// basics()

	vehicle2 := structs()
	fmt.Println(vehicle2)
}
