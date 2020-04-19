package main

// to run: 
// in /home/ericm/github/goapps/intro/code
// go run github.com/emacadie/goapps/intro/code
// or go run .
// you can do that in a directory with a mod w/a main
// or go build github.com/emacadie/goapps/intro/code
// it makes a new file

import (
	"fmt"
	"net/http"
	"github.com/emacadie/goapps/intro/code/controllers"
	"github.com/emacadie/goapps/intro/code/models"
)

func main() {
	fmt.Println("Ehllo from a module")
	u := models.User{
		ID: 2,
		FirstName: "George",
		LastName: "Washington",
	}
	fmt.Println("u: ", u)
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

