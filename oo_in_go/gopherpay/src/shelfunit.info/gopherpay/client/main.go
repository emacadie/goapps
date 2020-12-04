package main

import (
	"fmt"
	"shelfunit.info/gopherpay/payment"
)

func main() {
	credit := payment.CreateCreditAccount(
		"George Washington", 
		"1111-2222-3333-4444",
		5,
		2021,
		123)
	fmt.Printf( "Owner name: %s \n", credit.OwnerName() )
	fmt.Printf( "Card number: %s \n", credit.CardNumber() )
	fmt.Println( "Trying to change card number" )
	err := credit.SetCardNumber( "invalid" )
	if err != nil {
		fmt.Printf( "That didn't work: %v \n", err )
	}
	fmt.Println( "Trying again to set card number w/better regex"  )
	err = credit.SetCardNumber( "2222-3333-4444-5555" )
	if err != nil {
		fmt.Printf( "That didn't work either: %v \n", err )
	} else {
		fmt.Printf( "New card number: %s \n", credit.CardNumber() )
	}
	fmt.Printf( "Available credit: %v \n", credit.AvailableCredit() )
	
	// error if we try next line: credit.ownerName undefined (cannot refer to unexported field or method ownerName)
	// fmt.Printf( "Owner name: %s \n", credit.ownerName )

	fmt.Println( "About to call ProcessPayment" )
	cash := payment.CreateCashAccount()
	cash.ProcessPayment( 32 )
	credit.ProcessPayment( 44 )

	fmt.Printf(  "\n\n-------------------------------------------\n\n")

	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"John Adams",
		"1111-2222-3333-4445",
		5,
		2021,
		123)
	option.ProcessPayment( 500 )
	
	option = payment.CreateCashAccount()
	option.ProcessPayment( 500 )
	// what if you want to implement more than one interface?


}

