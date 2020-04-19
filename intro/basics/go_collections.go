package main

// to run on command line: go run go_collections.go 

import (
	"fmt"
)

func main() {
	// array: fixed size, same type
	var arr [3]int
	arr[0] = 1 // 0-based
	arr[1] = 2
	arr[2] = 3
	fmt.Println("arr: ", arr)
	fmt.Println("arr[1]: ", arr[1])

	// implicit initialization
	arr2 := [3]int{1,2,3}
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr2[1]: ", arr2[1])
	// arr3 := [3]int // this won't work
	arr2[1] = 5
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr2[1]: ", arr2[1])

	// slices: built on arrays
	// build slice from orig arr
	slice := arr[:] // go from beginning to end
	fmt.Println("arr: ", arr, ", slice: ", slice)
	arr[1] = 42
	slice[2] = 27
	fmt.Println("after change: arr: ", arr, ", slice: ", slice)
	// new slice
	slice2 := []int{1,2,3}
	fmt.Println("slice2: ", slice2)
	slice2 = append(slice2, 4) // sort of like functional langs
	fmt.Println("after adding to slice2: ", slice2)
	slice2 = append(slice2, 5, 6, 99)
	fmt.Println("after adding multiple elements to slice2: ", slice2)
	
	// a true slice:
	slice3 := slice2[1:]
	slice4 := slice2[:2]
	slice5 := slice2[1:2]
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
	fmt.Println("slice5: ", slice5)

	// maps
	// map with string key and int values
	m := map[string]int{"foo":42}
	fmt.Println("our map: ", m, "m[\"foo\"]:", m["foo"])
	m["foo"] = 27
	fmt.Println("our map after changing: ", m, "m[\"foo\"]:", m["foo"])
	// add a new key/value pair
	m["hello"] = 33
	fmt.Println("our map after adding: ", m, "m[\"foo\"]:", m["foo"])
	// delete a key and value
	delete(m, "foo")
	fmt.Println("our map after deleting: ", m, "m[\"foo\"]:", m["foo"])

	// struct - only collection allowing disparate data types
	// 2-step process
	type user struct {
		ID int
		FirstName string
		LastName string
	}
	// initialize it
	var u user
	fmt.Println("u:", u) // fields initialized to "zero value" or default
	u.ID = 1
	u.FirstName = "George"
	u.LastName = "Washington"
	fmt.Println("u after adding data:", u, ", First name:", u.FirstName)
	// shorter init syntax
	u2 := user{ID:2, FirstName: "John", LastName: "Adams"}
	fmt.Println("u2:", u2) 
	// if it is a big struct, you can init on multi lines
	// compiler wants you to end with a comma, so it is consistent with other lines
	u3 := user{ ID: 3,
		FirstName: "Thomas",
		LastName: "Jefferson",
	}
	fmt.Println("u3 with multi-line init:", u3) 
	

}

