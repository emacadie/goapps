package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println( "Here is the 1st time: ", t )
	time.Sleep( 2 * time.Second )
	t2 := time.Now()
	fmt.Println( "Here is the 2nd time: ", t2 )

	time.Sleep( 2 * time.Second )
	t3 := time.Now()
	fmt.Println( "Here is the 3rd time: ", t3 )

	Year := t3.Year()
	Month := t3.Month()
	Day := t3.Day()

	fmt.Printf( "Today is %d-%2d-%d \n", Year, Month, Day )
	fmt.Printf( "Time in ANSIC: %s \n", t3.Format( time.ANSIC ) )
	fmt.Printf( "I prefer RFC3339: %s \n", t3.Format( time.RFC3339 ) )
	fmt.Printf( "About to try custom formatting using reference date %s \n", "Mon Jan 2 15:04:05 MST 2006" )
	var customFormat string = "2006-01-02_15:04:05"
	fmt.Printf( "Time in custom format: %s \n", t3.Format( customFormat ) )

	startDate := time.Date( 
		2018, // year int, 
		05,   // month Month, 
		14,   // day, 
		10, // hour, 
		15, // min, 
		00, // sec, 
		00, // nsec int, 
		time.UTC ) // loc *Location)
	fmt.Println( "Here is startDate: ", startDate )
	fmt.Println( "Start date with custom formatting: ", startDate.Format( customFormat ) )
}

