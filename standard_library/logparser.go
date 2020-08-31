package main

// cd into logs, and run:
// go run secondlog.go
// get log.002.txt, and cp it into this dir

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len( os.Args ) > 1 {
		searchTerm := os.Args[ 1 ]
		file, _ := os.Open( "log.002.txt" )
		scanner := bufio.NewScanner( file )
		for scanner.Scan() {
			line := scanner.Text()
			result := strings.Contains( line, searchTerm )
			if result {
				fmt.Println( "Match: ", line)
			}
		} // for scanner.Scan() 
	} else {
		fmt.Println( "You must enter a search term" )
	} // if len( os.Args ) > 1 / else
}

