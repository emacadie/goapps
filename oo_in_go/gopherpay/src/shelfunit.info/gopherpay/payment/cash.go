package payment

import (
	"fmt"
)

// made for Encapsulation module

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ProcessPayment(amountA float32) bool {
	fmt.Printf("in ProcessPayment for cash with amount %.2f \n", amountA )
	return true
}


