package main

import (
	// "shelfunit.info/gopherpay/payment"
	"shelfunit.info/gopherpay/inherit"
	"shelfunit.info/gopherpay/paybroker"
)


type PaymentOption interface {
	ProcessPayment( float32 ) bool
}

func main() {
	// old way, before moving PaymentOption to here: var option inherit.PaymentOption
	var option PaymentOption
	

	option = &inherit.CreditCard {}
	option.ProcessPayment( 500 )

	option = &inherit.CheckingAccount {}
	option.ProcessPayment( 200 )

	option = &paybroker.PaymentBrokerAccount {}
	option.ProcessPayment( 250 )
}

