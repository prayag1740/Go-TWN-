//helper functions of main application
// a package is a collection of go files
//variables and functions outside any function can be accessed in all other files within the same package

package helper

import "strings"

//1st letter capitalize so that to make it available for all packages in the app
func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	var isValidTicket = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}

//we can also export variables in diff packages if we capitalize it ; GLOBAL variables
