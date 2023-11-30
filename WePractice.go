// This is strictly for practice
package main

import "fmt"

func main() {
	var Name string
	var Age int
	var Email string
	var TicketONo = 20
	var TicketNo int
	//To ask the user his name and age
	fmt.Printf("\nPlease provide us with your details")
	fmt.Scan(&Name, &Age, &Email, &TicketNo)
	TicketONo = TicketONo - TicketNo
	fmt.Printf("\nYour name is %v and your age is %v. Congratulations, you will recieve a confirmation mail in you %v showing your %v", Name, Age, Email, TicketONo)
}
