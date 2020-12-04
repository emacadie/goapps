package main

import (
	"fmt"
	"strings"
	"shelfunit.info/gopherpay/payment/objectoriented"
)

func main() {
	const amount = 500
	fmt.Println("Paying with cash")
	cash := &objectoriented.Cash{}
	cash.ProcessPayment( amount )
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
	
	credit := objectoriented.CreateCreditAccount(
		"George Washington",
		"1111-2222-3333-4444",
		5,
		2021,
		123) // go does not like having the close parens on its own line
 
	fmt.Println( "Paying with credit card" )
	fmt.Printf("Initial balance: $%.2f \n", credit.AvailableCredit())
	credit.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f \n", credit.AvailableCredit())
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := objectoriented.CreateCheckingAccount(
		"George Washington",
		"111111",
		"222222")

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f \n", checking.Balance())
	checking.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f \n", checking.Balance())
	fmt.Println("Not enough in the account, nothing we can do now")
	
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
	
}

