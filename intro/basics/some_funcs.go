package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello go")
	port := 3000
	err := startWebServer(port, 2)
	fmt.Println("err: ", err)
	// multiple return values
	returnPort, err2 := startWeb2(3001, 3)
	fmt.Println("returnPort: ", returnPort, ", err2: ", err2)
	// right-only variable if you only care about the second return value
	_, err3 := startWeb2(3001, 4)
	fmt.Println("err3: ", err3)
}

// you could also do this since we have 2 ints:
// func startWebServer( portArg, numRetries int)
func startWebServer(portArg int, numRetries int) error {
	fmt.Println("Starting func startWebServer")
	fmt.Println("Our port is:", portArg)
	fmt.Println("Number of retries: ", numRetries)
	// we don't throw exceptions, we return errors
	return errors.New("Something went wrong")
}

func startWeb2(portArg int, numRetries int) (int, error) {
	fmt.Println("Starting func startWeb2")
	fmt.Println("Our port is:", portArg)
	fmt.Println("Number of retries: ", numRetries)
	// we don't throw exceptions, we return errors
	return portArg, errors.New("Something went wrong")
}

