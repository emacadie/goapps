package objectoriented

// lower case names in struct means they cannot be accessed directly
// internal to package
type CheckingAccount struct {
	accountOwner  string
	routingNumber string
	accountNumber string
	balance       float32
}

// constructor function
func CreateCheckingAccount(accountOwner, routingNumber, accountNumber string ) *CheckingAccount {
	return &CheckingAccount {
		accountOwner:  accountOwner,
		routingNumber: routingNumber, 
		accountNumber: accountNumber,
		balance:       250, // this will come from a web service       
	}
}

func (c CheckingAccount) ProcessPayment(amount float32) bool {
	return true
}


func (c CheckingAccount) Balance() float32{
	return c.balance
}


