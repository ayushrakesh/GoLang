package main

import (
	"errors"
	"fmt"
)

type user struct {
	name        string
	number      int
	isScheduled bool
}

func createUserMap(names []string, phones []int) (map[string]int, error) {
	m := map[string]int{}

	// fmt.Println(m)
	// fmt.Println(len(m))

	if len(names) != len(phones) {
		return nil, errors.New("invalid sizes!")
	} else {
		for i := 0; i < len(names); i++ {
			m[names[i]] = phones[i]
		}
	}
	return m, nil
}
func deleteIfNecessary(userMap map[string]user, name string) (bool, error) {

	user, ok := userMap[name]
	if !ok {
		return false, errors.New("not found")
	} else {
		// user := userMap[name]

		if user.isScheduled {
			delete(userMap, name)
			return true, nil
		} else {
			return false, nil
		}
	}
}
func getMessageCounts(userIds []string) map[string]int {
	counts := make(map[string]int)

	for _, userId := range userIds {
		cnt := 0

		for j := 0; j < len(userIds); j++ {
			if userId == userIds[j] {
				cnt++
			}
		}
		counts[userId] = cnt
	}

	return counts
}
func main() {
	// m := map[string]int{
	// 	"john":     1,
	// 	"vicky":    2,
	// 	"steve":    3,
	// 	"mitchell": 4,
	// }

	names := []string{"john", "vicky", "mitchell", "steve"}
	phones := []int{101, 102, 103}

	m, err := createUserMap(names, phones)

	// var n map[string]int
	n := make(map[string]int) //nitilise empty array
	fmt.Println(m, err)
	fmt.Println(n)

	fmt.Println("-------------")

	userMap := map[string]user{
		"john": {
			name: "john", number: 123, isScheduled: true,
		},
		"vicky": {
			name:   "vicky",
			number: 124, isScheduled: false,
		},
		"steve": {
			name: "steve", number: 125, isScheduled: true,
		},
		"mitchell": {
			name: "mitchell", number: 126, isScheduled: true,
		},
	}

	isdeleted, err := deleteIfNecessary(userMap, "johni")

	fmt.Println(isdeleted, err, userMap)

	userIds := []string{
		"123", "123", "145", "145", "123", "145", "469",
	}

	l := getMessageCounts(userIds)

	fmt.Println(l)
}
