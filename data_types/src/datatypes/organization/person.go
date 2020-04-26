package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


// type alias
// we cannot make new funcs for it, since its a string
// type TwitterHandler = string

// type declaration
// this copies fields, alias copies fields and methods
// type TwitterHandler string
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf( "https://www.twitter.com/%s", cleanHandler )
}

type Identifiable interface {
	ID() string // it has one method that returns a string
}

type Citizen interface {
	Identifiable
	Country() string
}


type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

// europeanUnionIdentifier
type europeanUnionIdentifier struct {
	id string
	country string
}
/*
// old func
func NewEuropeanUnionIdentifier(idArg, countryArg string) Citizen {
	return europeanUnionIdentifier{
		id: idArg,
		// now he has this. when did that change?
		// country: []string{countryArg},
		country: countryArg,
	}
}
*/
// with the braces, the type could be anything
func NewEuropeanUnionIdentifier(idArg interface{}, countryArg string) Citizen {
	// we can switch on the type
	switch v := idArg.(type) { // go will case idArg to whatever it is
		case string:
			return europeanUnionIdentifier{
				id: v,
				country: countryArg,
			}
		case int:
		    return europeanUnionIdentifier{
				id: strconv.Itoa(v),
				country: countryArg,
			}
		case europeanUnionIdentifier:
		    return v
		/*
	    case Person:
		    euID, ok := v.Citizen(europeanUnionIdentifier)
		    return euID
		*/
		default:
		    panic("using invalid type for EU identifier")
	}

}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Employee struct {
	Name
}

type Person struct {
	// if lower-case, they are no longer public
	Name
	twitterHandler TwitterHandler
	// Identifiable // embedding interface
	Citizen // added after we embedded Identifiable in Citizen
}

// instead of Person AND Employee each having both firstName and lastName, let's make a Name struct
type Name struct {
	first string
	last string
}

// we can still call p.FullName() for Person
func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

// sort of like a constructor
func NewPerson(firstNameArg, lastNameArg string, citizen Citizen) Person {
	return Person {
		Name: Name {
			first: firstNameArg,
			last: lastNameArg,		
		},
		Citizen: citizen,
		// this was from when we had these as fields
		// firstName: firstNameArg,
		// lastName: lastNameArg,		
	}
}

// when you call, a copy is being generated, you are not changing state
// so make it a pointer
// you should make them all pointer-based to be consistent
// and it might be a large object
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

/*
no longer needed now that we have Name embedded in Person
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}
*/
// now Person implements the Identifiable interface
// so now we don't need this
/* 
func (p Person) ID() string {
	return "12345"
}
*/

