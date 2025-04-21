package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string // slice to hold bookings - a dynamic array

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	if userTickets > 50 {
		fmt.Printf("We only have %v tickets remaining, but you tried to book %v\n", remainingTickets+userTickets, userTickets)
		fmt.Printf("Please try again and book only %v tickets\n", remainingTickets+userTickets)
		return
	}
	if userTickets <= 0 {
		fmt.Printf("You need to book at least 1 ticket\n")
		return
	}

	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole bookings slice: %v\n", bookings)
	// fmt.Printf("The first booking is %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)
}
