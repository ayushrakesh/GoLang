// package main

// import "fmt"

// type email struct {
// 	body         string
// 	isSubscribed bool
// }

// type expense interface {
// 	cost() float64
// }
// type printer interface {
// 	print()
// }

// func (e email) cost() float64 {
// 	if e.isSubscribed {
// 		return 0.1 * float64(len(e.body))
// 	} else {
// 		return 0.5 * float64(len(e.body))
// 	}
// }
// func (e email) print() {
// 	fmt.Println(e.body)
// }
// func getDetails(e expense, p printer) string {
// 	p.print()
// 	return fmt.Sprintf("Expense is %v", e.cost())
// }
// func main() {
// 	s := email{
// 		body:         "test@gmail.com",
// 		isSubscribed: true,
// 	}

// 	fmt.Println(getDetails(s, s))
