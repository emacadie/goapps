package main

// you can run it like this:
//  go run showmestuff.go

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println( "Our current version of Go is: " + runtime.Version() )
	fmt.Printf( "w/printf: version is %v\n", runtime.Version() )
	fmt.Printf( "Using Go %v running in %v\n", runtime.Version(), runtime.GOOS )

	reader := bufio.NewReader( os.Stdin )
	fmt.Println( "What is your name: " )
	// https://golang.org/pkg/bufio/#Reader.ReadString
	text, _ := reader.ReadString( '\n' ) // this will keep return at the end
	fmt.Printf( "Hello %v \n", text )

}

