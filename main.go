package main

import "fmt"

func main() {

	var conferenceName = "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // specify data type if needed
	//floats for decimals

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conference name is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We currently have %v guests who booked tickets and %v spots remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend this")
	//inform when all 50 tickets are sold

	var givenName string
	var familyName string
	var email string
	var purchasedTickets int

	//ask user for name

	fmt.Println("Please enter your first name.")
	fmt.Scan(&givenName)

	fmt.Println("Please enter your last name.")
	fmt.Scan(&familyName)

	fmt.Println("Please enter your email.")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want?")
	fmt.Scan(&purchasedTickets)

	fmt.Printf("Hi, %v %v, thank you for your purchase of %v tickets.\nPlease check your email you registered at %v", givenName, familyName, purchasedTickets, email)

}
