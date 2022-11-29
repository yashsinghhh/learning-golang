package main

import (
	"fmt"
	"strings"
)

// we have to first initiate a go file that means that we
// we have to write a command in the terminal "go mod init "name of the folder"

// then also before the commands start running we have to define the package with
// the " package main " command

// we also have to define the entry point of the execution which is " func main() {    }"



	// we have to declare the package from which we are importing the commands so here we will import
	// the fmt package with "import "fmt" "

	//  now to know that what function is in which package we have to read dopcumentations or we can
	// simply google but wi`th practice we will be fluent with the packages

		//  printf can use the format to print the special charecter the format for it is as follows 
	

		// we have to define the values fo the variables but as most of the common variables are identified by the GO lang 
// it is only neccesary to define not so specific data types 
// such as uint where we will not be able to assign any negative value to the variable


// so var x int=1 can also be written as x:=1 
// it will define it as a variable and assign the specific variabel to it 


// fmt.Scanln(&userName)

// so in go when we ask to assign a value to the variable we have to define the specific memory address where 
//  GO will look for the variable THis is done by POinters

// CREAting an ARRAY 
//  var bookings [50]string
//  bookings[0]= firstname+" "+lastname 
//  bookings[1]= 

// fmt.Printf("the array %v is the booking list\n",bookings)
// fmt.Printf("the type of the bookings list is %T\n", bookings)
// fmt.Printf("the length of this array is %v\n", len(bookings))


// when using an array we have to define the size of the array but what if we do not have a definite value for it then 
// we use slice which is an abstraction og an array  

func main() {
	var conferenceName = "Go Conference"
	fmt.Println("Welcome to our booking appliation")
	fmt.Println("kindly follow the steps to book your ticket")

	fmt.Println(conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application")
	const conferenceTickets = 50
	var RemainingTickets = 50
	fmt.Println("Total tickets: ", conferenceTickets)
	fmt.Println("there are ", RemainingTickets, "tickets available")

	

fmt.Printf("there are %v tickets available\n", RemainingTickets)

var userName string
// asking for user name 




var userTickets int
userTickets=2
fmt.Printf("user %v is going to but %v tickets\n", userName, userTickets)
fmt.Printf("conferenceTickets is %T, conferconferenceName is %T\n",conferenceTickets, conferenceName)


tickets:= 4
fmt.Println(tickets)

// fmt.Scanln(&userName)



fmt.Println(userName)
fmt.Print("write name")
var firstname string
var lastname string
fmt.Scanln(&firstname)
fmt.Scanln(&lastname)






var bookings []string //this is an slice just a array with empty index bracket
// to add elements to the slice we can use 
bookings = append(bookings, firstname+" ",lastname+",")
fmt.Println(bookings)

// we can also do bookings:= []string{}

// FOR loop 
 for { 
	// var test1 string
	// var test2 string 
	// fmt.Println("first name ")
	// fmt.Scan(&test1)
	// fmt.Println("last name ")
	// fmt.Scan(&test2)
	// bookings = append(bookings,test1 + " " + test2+"," )
	// fmt.Println(bookings)

	fmt.Println(userName)
fmt.Print("write name")
var firstname string
var lastname string
fmt.Scanln(&firstname)
fmt.Scanln(&lastname)






var bookings []string //this is an slice just a array with empty index bracket
// to add elements to the slice we can use 
bookings = append(bookings, firstname+" ",lastname+",")


	firstNames:= []string{}
	for _, Name1:= range bookings {
	   var names = strings.Fields(Name1)
	   firstNames = append(firstNames, names[0])
	   
	   
}
fmt.Printf("first names list is %v",firstNames )

// for each loop 
// THis is used for iterating the list so if we just want the first name from the list we can do that 
// using for-each loop 
//  firstNames:= []string{}
//  for _, Name1:= range bookings {
// 	var names = strings.Fields(Name1)
// 	firstNames = append(firstNames, names[0])

	// as sometimes there are some unwanted variables that arent used so _ is used to identify them

	
 }
//  fmt.Printf("These are all the bookings: %v\n", bookings)

// range allows us to iterate through different datastructures not only list or slices 



}
