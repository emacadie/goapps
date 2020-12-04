package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// created for Encapsulation module

type PaymentOption interface {
	ProcessPayment( float32 ) bool
}

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}
// we cannot build one outside of package, so must provide constructor
// returning pointer
func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// we implement the interface by just implementing the method
func (c* CreditCard) ProcessPayment( amount float32 ) bool {
	fmt.Printf( "In ProcessPayment for a credit card with amount %.2f \n", amount )
	return true
}

// how can consumers interact with it?
// getters and setters
// look at function/data structure courses
// We have "(c CreditCard) right after "func"
func (c CreditCard) OwnerName() string {
	return c.ownerName
}
// this has pointer since we are changing data
func (c *CreditCard) SetOwnerName( value string) error {
	if len( value ) == 0 {
		return errors.New("Invalid owner name provided")
	}
	c.ownerName = value
	return nil
}
// why doesn't constructor have any validation like setters?
	
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

func (c *CreditCard) SetCardNumber( value string ) error {
	if !cardNumberPattern.Match( []byte( value ) ) {
		return errors.New( "Invalid credit card number format" )
	}
	c.cardNumber = value
	return nil
}

// this feels like a lot of boilerplate
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// but we can use one method each for two fields
func (c *CreditCard) SetExpirationDate( month, year int ) error {
	now := time.Now()
	if ( year < now.Year() ) || ( ( year == now.Year() ) && ( time.Month(month) < now.Month() ) ) {
		return errors.New( "Expiration date must be in the future" )
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetSecurityCode( value int ) {
	if value > 100 || value > 999 {
	}
}

func (c CreditCard) AvailableCredit() float32 {
	return 5000. // this could come from a web service
}




