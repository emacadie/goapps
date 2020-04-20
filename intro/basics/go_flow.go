package main

type User struct {
	ID int
	FirstName string
	LastName string
}

type HTTPRequest struct {
	Method string
}

func main() {
	// for loops
	// loop until condition
	var i int
	for i < 5 {
		println("loop until condition, i: ", i)
		i++
	}
	// break out early
	var i2 int
	for i2 < 5 {
		println("loop until condition w/break, i2: ", i2)
		i2++
		if i2 == 3 {
			break
		}
	}
	// continue
	var i3 int
	for i3 < 5 {
		i3++
		if i3 == 3 {
			continue
		}
		println("loop until condition w/continue, i3: ", i3)
	}

	// loop until condition w/post clause
	for i4 := 0; i4 < 5; i4++ {
		println("loop until condition w/post clause, i4: ", i4)
	}
	// cannot print i4 out here

	// infinite loop
	var i5 int
	for ; ; {
		println("infinite loop until condition w/post clause, i5: ", i5)
		if i5 == 5 {
			break
		}
		i5++
	}
	// better option
	var i6 int
	for {
		println("infinite loop until condition w/post clause, i6: ", i6)
		if i6 == 5 {
			break
		}
		i6++
	}
	
	// loop over collections
	slice := []int{1,2,3}
	for i7 := 0; i7 < len(slice); i7++ {
		println("next mem of slice: ", slice[i7])
	} 
	for i8, v8 := range slice {
		println("next mem of slice w/range: i8: ", i8, ", v8:", v8)
	} 
	// now with map
	wellKnownPorts := map[string]int {"http": 80, "https": 443}
	for k9, v9 := range wellKnownPorts {
		println("looping through map with key: ", k9, ", and value: ", v9)
	}
	// what if you just need to loop over keys?
	for k10 := range wellKnownPorts {
		println("looping through map with key only: ", k10)
	}
	// value only
	for _, v10 := range wellKnownPorts {
		println("looping through map with values only: ", v10)
	}

	// branching w/panic
	// panic("something bad just happened") // this will stop program from continuing
	println("After panic")
		
	u1 := User{ID: 1, FirstName: "George", LastName: "Washington"}
	u2 := User{ID: 2, FirstName: "John", LastName: "Adams"}
	// if
	if u1.ID == u2.ID {
		println("They are the same user")
	} else {
		println("dif user")
	}
	// could do more than one test - you can write branches where nothing prints
	if u1.ID == u2.ID {
		println("They are teh same user")
	} else if u1.FirstName == u2.FirstName {
		println( "same first name")
	} else {
		println("Still diff user")
	}
	
	// switch
	// no need for a break statement
	r := HTTPRequest{Method: "GET"}
	switch r.Method {
	case "GET":
		println("Get request")
	case "PUT":
		println("Put request")
	}
	// you can use "fallthrough" to go to next case
	switch r.Method {
	case "GET":
		println("Get request w/fallthrough")
		fallthrough
	case "POST":
		println("Post request after fallthrough")
	case "PUT":
		println("Put request")
	}
	// default
	r2 := HTTPRequest{Method: "NONE"}
	switch r2.Method {
	case "GET":
		println("Get request ")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Default")
	}
}

