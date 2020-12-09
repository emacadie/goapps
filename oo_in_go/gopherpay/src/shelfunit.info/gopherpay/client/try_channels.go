package main

import (
	"fmt"
)

type CreditAccount struct {}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Printf( "Processing credit card payment with amount: %.2f \n ", amount )
}
// constructor that takes a channel
func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	fmt.Println( "starting CreateCreditAccount" )
     creditAccount := &CreditAccount{}
     // using same channel
     go func(chargeCh chan float32) {
        for amount := range chargeCh {
            creditAccount.processPayment(amount)
        }
     } (chargeCh)
	fmt.Println( "About to return creditAccount from CreateCreditAccount" )
	return creditAccount
}

func main() {
	chargeCh := make( chan float32 )
	CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	// this is so channel gets the message
	fmt.Scanln( &a )

}

