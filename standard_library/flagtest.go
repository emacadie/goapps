package main

import (
	"flag"
	"fmt"
)

func main() {
	// command (argument) (argument)
	// command -flag
	// to get help, type "command -h"
	// https://golang.org/pkg/flag/#String
	// name, default value, text to describe flag
	archPtr := flag.String( "arch", "x86", "CPU Type" )
	

	osFlag := flag.String( "OS", "Linux", "Operating System" )
	
	flag.Parse() // must come after all flags
	// need to de-reference pointer
	switch *archPtr {
		case "x86":
		    fmt.Println( "arch: Running in 32 bit mode" )
		case "AMD64":
		    fmt.Println( "arch: Running in 64 bit mode" )
		    fmt.Println( "arch: It's what we all use" )
		case "IA64":
		    fmt.Println( "arch: Love Live Itanic" )
	}

	switch *osFlag {
		case "Linux":
		    fmt.Println( "OS: You are in the club" )
		case "FreeBSD":
		    fmt.Println( "OS: Truly you are a master" )
		case "Apple":
		    fmt.Println( "OS: If your friends jumped off a bridge, face it, so would you" )
	}

	fmt.Println( "Process complete" )
	 
}

