package inherit

import (
	"fmt"
)
/*
// interitance
type PaymentOption interface {
	ProcessPayment( float32 ) bool
}
*/
type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Printf( "In CreditCard.ProcessPayment with amount %.2f \n", amount )
	return true
}

type CheckingAccount struct{}

func (c *CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Printf( "In CheckingAccount.ProcessPayment with amount %.2f \n", amount )
	return true
}

