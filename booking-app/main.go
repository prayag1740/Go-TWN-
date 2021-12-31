// standard file where go code is executed
// All the go code needs to be in package

package main

import (
	"booking-app/helper" //our own package defined
	"fmt"                // fmt is the function import from main package
	"sync"
	"time"
)

var conferenceName = "Go Conference" //camel case for variables
//Pacakage level variables can be accessed between diff functions

// var bookings = make([]map[string]string, 1) //initial size   ---> in case of map
type UserData struct { //For saving different datatypes for an entity ; lightweight class
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var bookings = make([]UserData, 1) // creating a empty list of Userdata struct

var remainingTickets = 50

var wg = sync.WaitGroup{}

// entry point of our go application
func main() {

	// variable declaration

	const conferenceTickets = 50 //constant whose value cannot be changed
	//same shortcut way of declaring variable without datatype ; not applicable for constants

	greetUsers(conferenceTickets, remainingTickets)

	fmt.Printf("conferenceName type is %T, remainingTickets is %T\n", conferenceName, remainingTickets)
	//All variables expect datatypes ;; basic ones -- String and Integer
	// when we declare and initialize a variable in same line; no need to provide the data type as in case of conferenceName variable

	//ARRAYS --> need to define size at declaration (fized size)
	// var bookings [50]string //To store who created the bookings -----------> ARRAY

	//Slices abstraction of arrays ; no need to define size at declarartion

	//creating infinte loop for aksing user again and again
	for {

		firstName, lastName, email, userTickets := getUserInput()
		//Validations
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
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
		getUserMapData(firstName, lastName, email, userTickets)
		bookTickets(userTickets, firstName, lastName, email)

		firstNames := FirstNames()
		fmt.Printf("The first name of bookings are : %v\n", firstNames)

		wg.Add(1)                                              //sets the number of threads the main thread should wait for ; for our case it is 1
		go sendTicket(userTickets, firstName, lastName, email) //creating a seperate thread for execution in parallel; does not block
		//when main thread is done is does not wait for other threads

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}

	}
	wg.Wait() //waits for all other threads to be completed before completeting main thread
}

func greetUsers(confTickets int, remTickets int) {
	//using print formatting , does not create new line automatically
	// %v is the default formattor, other options also available
	fmt.Printf("Welcome to our conference booking:  %v\n", conferenceName)
	fmt.Println("We have a total of", confTickets, "tickets and", remTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func FirstNames() []string { //2nd value is return type
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

func bookTickets(userTickets int, fname string, lname string, email string) {
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fname, lname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func getUserMapData(firstName string, lastName string, email string, userTickets int) {

	//create a map for user (MAP)
	//in a map there can be no mix of data types
	// var userData = make(map[string]string) //empty map

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTicketsCount"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //saves in variable formatted string
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v to %v\n", ticket, email)
	fmt.Println("######################")
	wg.Done() //tells that is done
}

// go run main.go -- To run the a go file
// go run . -- to run all the files within current directory
