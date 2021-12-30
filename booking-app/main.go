// standard file where go code is executed
// All the go code needs to be in package

package main

import (
	"fmt" // fmt is the function import from main package
	"strings"
)

var conferenceName = "Go Conference" //camel case for variables
//Pacakage level variables can be accessed between diff functions

// entry point of our go application
func main() {

	// variable declaration

	const conferenceTickets = 50          //constant whose value cannot be changed
	remainingTickets := conferenceTickets //same shortcut way of declaring variable without datatype ; not applicable for constants

	greetUsers(conferenceTickets, remainingTickets)

	fmt.Printf("conferenceName type is %T, remainingTickets is %T\n", conferenceName, remainingTickets)
	//All variables expect datatypes ;; basic ones -- String and Integer
	// when we declare and initialize a variable in same line; no need to provide the data type as in case of conferenceName variable

	//ARRAYS --> need to define size at declaration (fized size)
	// var bookings [50]string //To store who created the bookings -----------> ARRAY

	//Slices abstraction of arrays ; no need to define size at declarartion
	var bookings []string

	//creating infinte loop for aksing user again and again
	for {

		firstName, lastName, email, userTickets := getUserInput()

		//Validations
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		var isValidTicket = userTickets > 0 && userTickets <= remainingTickets

		if !isValidName || !isValidEmail || !isValidTicket {
			if !isValidName {
				fmt.Println("first name and last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email Address you entered does not contain @ sign")
			}
			if !isValidTicket {
				fmt.Println("No of tickets you entered is invalid")
			}
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining . So you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		// bookings[0] = firstName + " " + lastName ----------> in case of arrays
		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets := bookTickets(remainingTickets, userTickets, firstName, lastName, email)

		firstNames := FirstNames(bookings)
		fmt.Printf("The first name of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}

	}

}

func greetUsers(confTickets int, remTickets int) {
	//using print formatting , does not create new line automatically
	// %v is the default formattor, other options also available
	fmt.Printf("Welcome to our conference booking:  %v\n", conferenceName)
	fmt.Println("We have a total of", confTickets, "tickets and", remTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func FirstNames(bookings []string) []string { //2nd value is return type
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) // splits the string with white space as seperator
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTickets int

	//Taking input from user
	fmt.Println("enter your first name : ")
	fmt.Scan(&firstName) //whatever user enters assign that value
	//a pointer is a variable that points to the memory address of another variable
	// assign value in username memory that is why pointer is used

	fmt.Println("enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email : ")
	fmt.Scan(&email)

	fmt.Println("enter no of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets int, userTickets int, fname string, lname string, email string) int {
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fname, lname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets

}

// go run main.go -- To run the a go file
