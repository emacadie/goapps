package main

// /home/ericm/Downloads/zip.files/go.lib/06/demos/bigcustomerlist.csv
// /home/ericm/Downloads/zip.files/go.lib/06/demos/customerlist.csv


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args
	// open the customer list
	custlist, err := os.Open( string( args[ 1 ] ) )
	check( err )
	defer custlist.Close()

	timeTrack( start, "Opened customer list" )

	// create an output file
	outfile, err := os.Create( "outfile.csv" )
	check( err )
	defer outfile.Close()
	
	timeTrack( start, "Created output file" )
	
	// scan the customer list
	scanner := bufio.NewScanner( custlist )
	for scanner.Scan() {
		names := strings.Split( scanner.Text(), "," )
		outfile.WriteString( names[ 1 ] + "," + names[ 2 ] + "\n" )
	}
	check( scanner.Err() )
	timeTrack( start, "Wrote data to outfile" )
	defer timeTrack( start, "App closes" )
}

func timeTrack( start time.Time, name string ) {
	elapsed := time.Since( start )
	log.Printf( "%s took %s", name, elapsed )
	fmt.Printf( "%s took %s \n", name, elapsed )
}

// He really should have shown this before running
func check( e error ) {
	if e != nil {
		log.Fatal( e )
	}
}


