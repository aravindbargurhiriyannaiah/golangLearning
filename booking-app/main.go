package main

import (
	"fmt"
	"sync"
	"time"
)

const totalTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	// for {
	firstName, lastName, emailAddress, noOfTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketCount := validateInput(firstName, noOfTickets, emailAddress, lastName, remainingTickets)

	if isValidTicketCount && isValidName && isValidEmail {
		bookTicket(noOfTickets, firstName, lastName, emailAddress)
		wg.Add(1)
		go sendTicket(noOfTickets, firstName, lastName, emailAddress)
		printFirstNames()
		noTicketsAvailable := remainingTickets == 0
		if noTicketsAvailable {
			fmt.Println("The conference is booked out. Please come back next year!")
		}
	} else {
		if !isValidName {
			fmt.Printf("First name, %v, and/or last name, %v, you entered is too short\n", firstName, lastName)
		} else if !isValidEmail {
			fmt.Printf("Email address you entered, %v, is not valid\n", emailAddress)
		} else if !isValidTicketCount {
			fmt.Printf("Invalid number of tickets to book. We only have %v tickets left.\n", remainingTickets)
		}
	}
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available\n", totalTickets, remainingTickets)
	fmt.Println("Please book your tickets here.")
}

func printFirstNames() {
	firstNamesSlice := []string{}
	for _, booking := range bookings {
		firstNamesSlice = append(firstNamesSlice, booking.firstName)
	}
	fmt.Printf("Total number of bookings = %v\nList of all bookings' first names are: %v\n", len(bookings), firstNamesSlice)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	fmt.Println("Please enter you first name:")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	var emailAddress string
	fmt.Println("Please enter your email address: ")
	fmt.Scan(&emailAddress)

	var noOfTickets uint
	fmt.Println("How many tickets do you want to book:")
	fmt.Scan(&noOfTickets)

	return firstName, lastName, emailAddress, noOfTickets
}

func bookTicket(noOfTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets -= noOfTickets
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       emailAddress,
		noOfTickets: noOfTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, lastName, noOfTickets, emailAddress)
	fmt.Printf("%v ticket remaining for %v!\n", remainingTickets, conferenceName)
	fmt.Printf("The total bookings list is %v\n", bookings)
}

func sendTicket(noOfTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	fmt.Println("---------------------")
	var ticket = fmt.Sprintf("%v tickets for %v %v sent to %v", noOfTickets, firstName, lastName, emailAddress)
	fmt.Println(ticket)
	fmt.Println("---------------------")
	wg.Done()
}
