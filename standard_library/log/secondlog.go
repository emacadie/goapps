package main

import (
	// "fmt"
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile( "log.002.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666 )
	if err != nil {
		log.Fatal( err )
	}

	InfoLogger    = log.New( file, "INFO: ",    log.Ldate|log.Ltime|log.Lshortfile )
	WarningLogger = log.New( file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile )
	ErrorLogger   = log.New( file, "ERROR: ",   log.Ldate|log.Ltime|log.Lshortfile )
	FatalLogger   = log.New( file, "FATAL: ",   log.Ldate|log.Ltime|log.Lshortfile )

}

func main() {
	InfoLogger.Println( "This is info" )
	WarningLogger.Println( "I'm warning you" )
	ErrorLogger.Println( "Error message" )
	FatalLogger.Fatal( "It's over" )
}


