package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	// p := organization.Person{}
	// println(p.ID())
	// we could also do this:
	// var p2 organization.Identifiable = organization.Person{}
	// here, since you are setting p2 as an Identifiable, you cannot access FirstName or LastName
	// which are defined on Person
	// if you try, you get this:
	// p2.FirstName undefined (type organization.Identifiable has no field or method FirstName)
	// println(p2.ID()) // this won't work since Person now has an Identifiable inside it
	// but we can for p
	// p.FirstName = "George"
	// fmt.Println("Here is p: ", p)
	var p3 = organization.NewPerson("John", "Adams", organization.NewSocialSecurityNumber("123-45-6789"))
	fmt.Println("Here is p3:", p3, ", full name:", p3.FullName())
	err := p3.SetTwitterHandler("prez2")
	if err != nil {
		fmt.Printf("Error occurred in setting twitter handler: %s \n", err.Error())
	}
	err2 := p3.SetTwitterHandler("@prez2")
	if err2 != nil {
		fmt.Printf("Error occurred in setting twitter handler: %s \n", err2.Error())
	} else {
		fmt.Println("Success setting twitterHandler: ", p3, ", ", p3.TwitterHandler())
	}
	fmt.Println("calling RedirectUrl on p3.TwitterHandler:", p3.TwitterHandler().RedirectUrl())
	// type aliases, what is type of twitter handler?
	fmt.Printf("Type of TwitterHandler: %T \n", organization.TwitterHandler("test"))
	fmt.Println("calling p3.ID(): ", p3.ID(), ", p3.Country(): ", p3.Country())


	var pJQ = organization.NewPerson("Jaques", "Chirac", organization.NewEuropeanUnionIdentifier("123-45-6789", "France"))
	fmt.Println("Here is pJQ:", pJQ)
	fmt.Println("Country of pJQ:", pJQ.Country())

	name1 := NameB{"George", "Washington"}
	name2 := NameB{"George", "Washington"}

	if name1 == name2 {
		println("names match")
	} else {
		println("names do not match")
	}
	otherName1 := OtherName{"John", "Adams"}
	fmt.Println("otherName1: ", otherName1)
	/*
    // we cannot do this even though the field names and types are the same
	if name1 == otherName1 {
		println("name1 matches otherName1")
	}
	*/

	// comparing interfaces
	ssn := organization.NewSocialSecurityNumber("234-56-7890")
	eui := organization.NewEuropeanUnionIdentifier("234-56-7890", "France")
	eui2 := organization.NewEuropeanUnionIdentifier("234-56-7890", "France")
	// we can compare since the "New" funcs return a Citizen interface, not underlying types
	if ssn == eui {
		println("ssn is equal to eui")
	} else {
		println("ssn is not equal to eui")
	}
	if eui == eui2 {
		println("eui is equal to eui2")
	}
	// but if eu had country as slices, this would not work

	// zero value comparison
	if name1 == (NameB{}) {
		println("name1 empty")
	} else {
		println("name1 not empty")
	}
	name3 := NameB{"", ""}
	if name3 == (NameB{}) {
		println("name3 empty")
	} else {
		println("name3 not empty")
	}
	// this is better than checking if it is nil
	// pointers get allocated to heap, zero value comparison uses stack

	// keys for maps
	portfolio := map[NameB][]organization.Person{}
	portfolio[name1] = []organization.Person{p3}
	// you can use a struct as a key if the sub-types are primitive, like int and string
	// strict memory layout - it gives you comparable and hashable
	// other wise, you could extract easy parts to a subtype,
	// or make your own Equals function
	// See NameC and the Equals for NameC

	// switching on types
	// id is a string
	var pJQ1 = organization.NewPerson("Jaques1", "Chirac1", organization.NewEuropeanUnionIdentifier("123-45-6789", "France"))
	// id is an int
	var pJQ2 = organization.NewPerson("Jaques2", "Chirac2", organization.NewEuropeanUnionIdentifier(123456789, "France"))
	fmt.Println("Here is pJQ1:", pJQ1)
	fmt.Println("Here is pJQ2:", pJQ2)
	
}

type NameB struct {
	First string
	Last string
	// Middle []string // go can only use == if all types in struct are primitive
}

// need to make your own Equals
func (n NameC) Equals (otherName NameC) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

type NameC struct {
	First string
	Last string
	Middle []string // go can only use == if all types in struct are primitive
}

type OtherName struct {
	First string
	Last string
}

