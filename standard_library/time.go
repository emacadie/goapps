package main

import(
	"fmt"
	"time"
)

func main() {
	var ourFormat string = "Monday, January 2, 2006"
	var minuteFormat string = "15:04:05"
	first := time.Now()
	fmt.Printf( "It is currently %v \n", first.Format( minuteFormat ) )
	time.Sleep( 2 * time.Second )
	second := time.Now()
	fmt.Printf( "It is now %v \n", second.Format( minuteFormat ) )
	
	// fmt.Printf( "About to try custom formatting using reference date %s \n", "Mon Jan 2 15:04:05 MST 2006" )   
	today := time.Now()
	fmt.Printf( "It is now %v \n", today.Format( ourFormat ) )
	// how much time has passed since 2018-07-04
	startDate := time.Date( 
		2018, // year int, 
		07,   // month Month, 
		04,   // day, 
		9, // hour, (Is he using a different function? why does he have a 9?)
		00, // min, 
		00, // sec, 
		00, // nsec int, 
		time.UTC ) // loc *Location)
	elapsed := time.Since( startDate )
	fmt.Printf( "%v has passed since %v \n", elapsed, startDate.Format( ourFormat ) )
	// gives something like 18764h55m53.301228724s - not super helpful
	fmt.Printf( "Hours: %v Minutes: %v Seconds: %v \n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds() )
	fmt.Printf( "Hours: %.0f Minutes: %.0f Seconds: %.0f \n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds() )
	// there are external libraries that do date manipulation

	// date in six months
	today2 := time.Now()
	future := today2.AddDate( 
		0,  // years int, 
		6,  // months int, 
		0 ) // days int
	fmt.Printf( "In six months it will be %v \n", future.Format( ourFormat ) )
	// to go backward, use negative number
	past := today2.AddDate( 
		0,   // years int, 
		-6,  // months int, 
		0 )  // days int
	fmt.Printf( "Six months ago it was %v \n", past.Format( ourFormat ) )

	sixHours := today2.Add( 6 * time.Hour )
	fmt.Printf( "In six hourss it will be %v \n", sixHours.Format( minuteFormat ) )
	// could not call variable "90Days", but "in90Days" worked, so I guess they cannot start with a number.
	in90Days := today2.AddDate( 
		0,   // years int, 
		0,   // months int, 
		90 ) // days int
	fmt.Printf( "In 90 days it will be %v \n", in90Days.Format( ourFormat ) )

	fmt.Println( "Let's set a deadline" )
	bedtime := time.Date(
		2020, // year int, 
		8,   // month Month, 
		24,   // day, 
		23,   // hour, (Is he using a different function? why does he have a 9?)
		00,   // min, 
		00,   // sec, 
		00,   // nsec int, 
		time.Local ) // loc *Location)
	fmt.Printf( "There are %0.f hours until bed time \n", time.Until( bedtime ).Hours() )
}


