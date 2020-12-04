package main

import (
	"fmt"
	"strings"
	"shelfunit.info/gopherpay/payment/procedural"
)

func main() {
	const amount = 500

	fmt.Println( "Paying with cash" )
	procedural.PayWithCash( amount )
	fmt.Printf( strings.Repeat( "*", 10 ) + "\n\n"  )
	
	credit := &procedural.CreditCard{
		OwnerName:       "George Washington",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 5,
		ExpirationYear:  2021,
		SecurityCode:    123,
		AvailableCredit: 5000, // don't forget comma at the end, should we be able to set this?
	}
	fmt.Println( "Paying with credit card" )
	fmt.Printf( "Initial balance: $%.2f \n", credit.AvailableCredit )
	procedural.PayWithCredit( credit, amount )
	fmt.Printf( "Balance now: $%.2f \n", credit.AvailableCredit )
	fmt.Printf( strings.Repeat( "*", 10 ) + "\n\n"  )

	checking := &procedural.CheckingAccount {
		AccountOwner:  "George Washington",
		RoutingNumber: "012345678",
		AccountNumber: "123456789",
		Balance:       250, // should we be able to set this?
	}

	fmt.Println( "Paying with check" )
	fmt.Printf( "Initial balance: $%.2f \n", checking.Balance )
	procedural.PayWithCheck( checking, amount )
	fmt.Printf( "Balance now: $%.2f \n", checking.Balance )
	
	fmt.Println( "Not enough. We can fix that." )
	checking.Balance = 5000

	fmt.Println( "Paying with check" )
	fmt.Printf( "Initial balance: $%.2f \n", checking.Balance )
	procedural.PayWithCheck( checking, amount )
	fmt.Printf( "Balance now: $%.2f \n", checking.Balance )


	fmt.Printf( strings.Repeat( "*", 10 ) + "\n\n"  )
	


} // end main


