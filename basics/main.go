package main

import (
	"fmt"
	"math"
)

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
type authenticationInfo struct {
	username string
	password string
}
type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius int
}
type rect struct {
	height int
	width  int
}

func getArea(s shape) float64 {
	return s.area()
}
func getPerimeter(s shape) float64 {
	return s.perimeter()
}
func (c circle) area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * float64(c.radius)
}
func (r rect) area() float64 {
	return float64(r.height) * float64(r.width)
}
func (r rect) perimeter() float64 {
	return 2 * float64((r.height + r.width))
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

type email struct {
	body         string
	isSubscribed bool
}

type expense interface {
	cost() float64
}
type printer interface {
	print()
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return 0.1 * float64(len(e.body))
	} else {
		return 0.5 * float64(len(e.body))
	}
}
func (e email) print() {
	fmt.Println(e.body)
}
func getDetails(e expense, p printer) string {
	p.print()
	return fmt.Sprintf("Expense is %v", e.cost())
}

func (ai authenticationInfo) userAuth() string {
	return fmt.Sprintf("Authentication %s+%s", ai.username, ai.password)
}

func main() {

	// basics()

	vehicle2 := structs()
	fmt.Println(vehicle2)

	user1 := authenticationInfo{
		username: "user-1",
		password: "123456",
	}

	fmt.Println(user1.userAuth())

	c := circle{
		radius: 3.0,
	}
	d := rect{
		height: 5,
		width:  10,
		// radius: 3.0,
	}
	fmt.Println(getArea(c), getArea(d))
	fmt.Println(getPerimeter(c), getPerimeter(d))

	s := email{
		body:         "test@gmail.com",
		isSubscribed: true,
	}

	fmt.Println(getDetails(s, s))
}
