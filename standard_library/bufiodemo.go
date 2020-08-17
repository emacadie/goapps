package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner( os.Stdin )
	/*
    // I hate it when they do different things in the same file
	fmt.Println( "Scanner will scan until you type \"/quit\"" )
	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println( "quitting" )
			os.Exit( 3 )
		} else {
			fmt.Println( "You typed: " + scanner.Text() )
		}
	} // for scanner.Scan() 
    */
	file, err := os.Open( "test.txt" )
	if err != nil {
		fmt.Println( err )
	}
	defer file.Close()
	fileScanner := bufio.NewScanner( file )
	for fileScanner.Scan() {
		fmt.Println( fileScanner.Text() )
	}
	fmt.Println( "Done with file" )


	if err := scanner.Err(); err != nil {
		fmt.Println( err )
	}

}

