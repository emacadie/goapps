package main

import (
	"fmt"
)

type CreditAccount struct {}

func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println( "In CreditAccount.AvailableFunds()" )
	return 250
}

type CheckingAccount struct {}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println( "In CheckingAccount.AvailableFunds()" )
	return 100
}

type HybridAccount struct {
	CreditAccount
	CheckingAccount
}

func (h *HybridAccount) AvailableFunds() float32 {
	fmt.Println( "In HybridAccount.AvailableFunds()" )
	// we could do any of these
	// return h.CreditAccount.AvailableFunds() + h.CheckingAccount.AvailableFunds()
	// return h.CreditAccount.AvailableFunds()
	// return h.CheckingAccount.AvailableFunds()
	return 500
}

func main() {
	ha := &HybridAccount{}
	// first time, before fix: ambiguous selector ha.AvailableFunds
	fmt.Println( ha.AvailableFunds() )

	// could also do this:
	fmt.Println( "About to explicitly call CheckingAccount.AvailableFunds through the HybridAccount" )
	fmt.Println( ha.CheckingAccount.AvailableFunds() )
}

