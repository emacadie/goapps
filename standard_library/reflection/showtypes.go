package main

import (
	"fmt"
	"reflect"
)

func main() {
	type person struct {
		personId  int
		firstName string
		lastName  string
	}

	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
		
	}

	newPerson := person{ 0, "George", "Washington" }
	fmt.Printf( "Our person is %s %s with an id of %d \n", newPerson.firstName, newPerson.lastName, newPerson.personId )

	// reflection
	fmt.Printf( "newPerson has type of %v \n", reflect.TypeOf( newPerson ) )
	fmt.Printf( "newPerson has value of %v \n", reflect.ValueOf( newPerson ) )
	fmt.Printf( "newPerson has a kind of %v \n", reflect.ValueOf( newPerson ).Kind() )

	fmt.Println()
	
	newEmployee := employee{ 0, "George", "Washington" }
	newCustomer := customer{ 0, "Peyton", "Randolph", "Virginia House of Burgesses" }
	addPerson( newEmployee )
	addPerson( newCustomer )
}

// function to add Employees or Customers to a database, without making separate functions
func addPerson( p interface{} ) bool {
	if reflect.ValueOf( p ).Kind() == reflect.Struct {
		fmt.Println( "In function addPerson: It's a struct" )
		v := reflect.ValueOf( p )
		switch reflect.TypeOf( p ).Name() {
			case "employee":
			    empSqlString := "Insert into employees( personId, firstName, lastName ) values ( ?, ?, ? )"
			    fmt.Printf( "SQL: %s \n", empSqlString )
			    fmt.Printf( "Added %v \n", v.Field( 1 ) )
		    case "customer":
			    custSqlString := "Insert into customers( personId, firstName, lastName, company ) values ( ?, ?, ?, ? )"
			    fmt.Printf( "SQL: %s \n", custSqlString )
			    fmt.Printf( "Added %v \n", v.Field( 1 ) )
		}
		return true
	} else {
		return false
	}
}

