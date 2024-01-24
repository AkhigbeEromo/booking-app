package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
	
		bookings := []string{}

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName:= len(firstName) >= 2 && len(lastName) >=2
		isValidEmail:= strings.Contains(email, "@")
		isValidTicketNumber:= userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{
			remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0], names[1])
		}
		fmt.Printf("The first names of booking are: %v\n", firstNames)

	

		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year")
			break
		}
			
		}else {
			fmt.Println("Your validation is wrong, redo it ")
		}
		
	}

}
