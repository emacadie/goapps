package paybroker

import (
	"fmt"
)

type PaymentBrokerAccount struct {}

// this is a PaymentOption instance
func (p *PaymentBrokerAccount) ProcessPayment(amount float32) bool {
	fmt.Printf( "In PaymentBrokerAccount.ProcessPayment with amount %.2f \n", amount )
	return true
}

