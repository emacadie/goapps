package main

import (
	"fmt"
	"os"
)

// go build dosomething.go
// args are 0-indexed
func main() {
	args := os.Args
	fmt.Println( args )
}

