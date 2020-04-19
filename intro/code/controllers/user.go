package controllers

import (
	"net/http"
	"regexp"
)

// use an empty struct to associate behaviors (sometimes, but not here)
type userController struct {
	userIDPattern *regexp.Regexp
}

// implementing net/http/Handler interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from user controller")) // convert a string to a slice of bytes
}

// constructor function
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile("^/users/(\\d+)/?"),
	}
}


