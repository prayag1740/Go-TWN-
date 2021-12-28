// standard file where go code is executed
// All the go code needs to be in package

package main

import "fmt" // fmt is the function import from main package

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

	var userName string
	var userTickets int

	//Taking input from user
	fmt.Println("enter your name : ")
	fmt.Scan(&userName) //whatever user enters assign that value
	//a pointer is a variable that points to the memory address of another variable
	// assign value in username memory that is why pointer is used

	userTickets = 2
	fmt.Printf("User %v booked %v tickets \n", userName, userTickets)

}

// go run main.go -- To run the a go file
