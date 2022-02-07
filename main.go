package main

import "fmt"

func main() {

	var conferenceName = "Golang Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We currently have %v guests who booked tickets and %v spots remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend this")

	//inform when all 50 tickets are sold

}
