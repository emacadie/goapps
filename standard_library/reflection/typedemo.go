package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	personId  int
	firstName string
	lastName  string
}

func main() {
	newEmployee := employee{ 0, "George", "Washington" }
	// kind is struct, type is employee, value is the values of the fields

	// name
	fmt.Printf( "newEmployee has a name of %v \n", reflect.TypeOf( newEmployee ).Name() )
	// type
	fmt.Printf( "newEmployee has type of %v \n", reflect.TypeOf( newEmployee ) )
	// kind
	fmt.Printf( "newEmployee has a kind of %v \n", reflect.TypeOf( newEmployee ).Kind() )
	// value
	fmt.Printf( "newEmployee has value of %v \n", reflect.ValueOf( newEmployee ) )

	fmt.Println()

	// slice
	employees := make([]employee, 3)
	employees = append( employees, employee{ 2, "John", "Adams" } )
	employees = append( employees, employee{ 3, "Thomas", "Jefferson" } )
	employees = append( employees, employee{ 4, "James", "Madison" } )
	// name of slice is blank
	fmt.Printf( "employees has a name of %v \n", reflect.TypeOf( employees ).Name() )
	// type
	fmt.Printf( "employees has type of %v \n", reflect.TypeOf( employees ) )
	// kind
	fmt.Printf( "employees has a kind of %v \n", reflect.TypeOf( employees ).Kind() )
	// value
	fmt.Printf( "employees has value of %v \n", reflect.ValueOf( employees ) )

	// make another slice with reflection
	eType := reflect.TypeOf( employees )
	newEmployeeList := reflect.MakeSlice( eType, 0, 0 )
	newEmployeeList = reflect.Append( newEmployeeList, reflect.ValueOf( employee{ 5, "James", "Monroe"} ) )
	fmt.Printf( "First list of employees: %v\n\n", employees )
	fmt.Printf( "List created by reflection: %v\n", newEmployeeList )

}

