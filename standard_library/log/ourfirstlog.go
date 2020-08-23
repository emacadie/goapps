package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func main() {
	log.Println( "This is our first message" )
	
	
	log.Println( "This is our second message" )
	writeLog( INFO, "Hello to our file" )
	writeLog( WARNING, "This is a warning" )
	writeLog( ERROR, "Getting worse" )
	writeLog( FATAL, "Damn fatal" )
}

func writeLog( mtype messageType, message string ) {
	// do we want to check for the log file every single time we call this fuction?
	file, err := os.OpenFile( "log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666 )
	if err != nil {
		log.Fatal( err )
	}
	
	log.SetOutput( file )

	switch mtype {
		case INFO:
		    logger := log.New( file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile )
		    logger.Println( message )
		case WARNING:
		    logger := log.New( file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile )
		    logger.Println( message )
		case ERROR:
		    logger := log.New( file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile )
		    logger.Println( message )
		case FATAL:
		    logger := log.New( file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile )
		    logger.Fatal( message ) // this will cause app to exit
	}

}


