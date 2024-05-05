package main

type msgSend struct {
	message   string
	sender    users
	recepient users
}

type users struct {
	name  string
	phone int
}
