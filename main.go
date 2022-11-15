package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		//keeping track of ticket sales
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		//This loop ends when iterated over all elements of the bookings list
		fmt.Printf("These first names of bookings are: %v\n", firstNames)
	}

	//var bookings = [50]string{}
	//variables & types
	// using an underscore _ is a blank Identifier it ignores a variable you don't want to use
	//bookings[1] = "Anna Lee"
	//RANGE - iterates over elements for different data structures(so not only arrays and slices)
	// For arrays & slices, RANGE provides the index and value for each element
	//string.Fields - Splits the string with white space as a separator and returns a slice with the split elements

}

//arrays & slices in Go

//Slice an abstraction of an ARRAY
//Slices are more flexible - variable-length or get a sub-array of its own
//Slices are also index-based and have a size, but is resized when needed.

//for loop 1:15:27
//https://www.youtube.com/watch?v=yyUHQIec83I&t=3078s

//infinite loop - keeping tab of people who book tickets as well as the number of tickets sold
