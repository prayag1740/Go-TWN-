// standard file where go code is executed
// All the go code needs to be in package

package main

import (
	"fmt" // fmt is the function import from main package
	"strings"
)

// entry point of our go application
func main() {

	// variable declaration
	var conferenceName = "Go Conference"  //camel case for variables
	const conferenceTickets = 50          //constant whose value cannot be changed
	remainingTickets := conferenceTickets //same shortcut way of declaring variable without datatype ; not applicable for constants

	fmt.Printf("conferenceName type is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName) //using print formatting , does not create new line automatically
	// %v is the default formattor, other options also available
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	//All variables expect datatypes ;; basic ones -- String and Integer
	// when we declare and initialize a variable in same line; no need to provide the data type as in case of conferenceName variable

	var firstName string
	var lastName string
	var email string
	var userTickets int

	//ARRAYS --> need to define size at declaration (fized size)
	// var bookings [50]string //To store who created the bookings -----------> ARRAY

	//Slices abstraction of arrays ; no need to define size at declarartion
	var bookings []string

	//creating infinte loop for aksing user again and again
	for {

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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining . So you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		// bookings[0] = firstName + " " + lastName ----------> in case of arrays
		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking) // splits the string with white space as seperator
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("The first name of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}

	}

}

// go run main.go -- To run the a go file
