package main

import (
	"fmt"
	"strings"
)

func main() { 
	bookings:= []string{}
	var remainingTickets = 5 
	var firstn = []string{}
	fmt.Print("tickets you want to book: ")
	var userTickets int
	fmt.Scan(&userTickets)
for{
	
	
	var firstname string
	var lastname string 
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstname)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastname)
	bookings = append(bookings, firstname+" "+lastname+",")
	if userTickets>remainingTickets {
		fmt.Printf("sorry, only %v tickets are available at the moment :(",remainingTickets)
		continue}
	remainingTickets= remainingTickets-1

	
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstn = append(firstn, names[0])
	}
	

	if remainingTickets==0{
		fmt.Print("fully booked","\n")
		break
	}

	
	


	
	
}

fmt.Print(bookings,"\n")




}