package main

import (
	"strings"
)

func validateInput(firstName string, noOfTickets uint, emailAddress string, lastName string, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@") && strings.Contains(emailAddress, ".")
	isValidTicketCount := noOfTickets > 0 && noOfTickets < remainingTickets
	return isValidName, isValidEmail, isValidTicketCount
}
