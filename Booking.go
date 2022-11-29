package main 
import ("fmt"
)
func main()   {
for { 
	
	var availabletickets int
	availabletickets=10
	var firstname string
	var lastname string
	fmt.Println("Greetings!")
	fmt.Print("First name: ")
	fmt.Scan(&firstname)
	fmt.Print("Last name: ")
	fmt.Scan(&lastname)
	fmt.Print("Enter the number of tickets you want to book: ")
	var ticketNumber int
	fmt.Scan(&ticketNumber)
	if ticketNumber<availabletickets {
		availabletickets=availabletickets-ticketNumber
		fmt.Printf("You have succesfully booked %v tickets!",ticketNumber)
	break
		} else { 
			fmt.Printf("sorry only %v tickets are available\n",availabletickets)
		continue}
	}

}
	
