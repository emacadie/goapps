package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ourString := "this is a string"
	stringCollection := strings.Split( ourString, " " )
	for i := range stringCollection {
		fmt.Println( "Next part of '", ourString, "': ", stringCollection[ i ] )
	}

	fmt.Println()

	ourString02 := "this is a string|this is another one|layin' some pipe"
	stringCollection02 := strings.Split( ourString02, "|" )
	for i := range stringCollection02 {
		fmt.Println( "Next part of '", ourString02, "': ", stringCollection02[ i ] )
	}

	fmt.Println()
	// to include the pipes, use SplitAfter
	stringCollection03 := strings.SplitAfter( ourString02, "|" )
	for i := range stringCollection03 {
		fmt.Println( "Next part of '", ourString02, "': ", stringCollection03[ i ] )
	}

	fmt.Println()
	// what if you only want the first 2 splits? strings.SplitAfterN
	// using our original string
	stringCollection04 := strings.SplitAfterN( ourString, " ", 2 )
	for i := range stringCollection04 {
		fmt.Println( "Next part of '", ourString, "': ", stringCollection04[ i ] )
	}

	fmt.Println()
	// splitting on new line
	ourString05 := "This is a string\nThis is the second string"
	stringCollection05 := strings.Split( ourString05, "\n" )
	for i := range stringCollection05 {
		fmt.Println( "Next part of '", ourString05, "': ", stringCollection05[ i ] )
	}

	fmt.Println( "\n-----\n" )
	// make sure you have a file
	file, _ := os.Open( "custlist.csv" )
	scanner := bufio.NewScanner( file )

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split( line, "," )
		fmt.Println( " -- new record --" )
		for i := range items {
			fmt.Println( items[ i ] )
		}
		fmt.Println()
	}
	


}

