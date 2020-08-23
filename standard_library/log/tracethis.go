package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"runtime/trace"
)

func main() {
	f, err := os.Create( "trace.out" )
	if err != nil {
		log.Fatalf( "We did not create trace file: %v \n", err )
	}
	// this is like an inline function
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf( "Unable to close trace file: %v \n", err )
		}
	}()

	if err := trace.Start( f ); err != nil {
		log.Fatalf( "Failed to start trace: %v \n", err )
	}
	defer trace.Stop()


	AddRandomNumbers()

}

func AddRandomNumbers() {
	firstNum  := rand.Intn( 100 )
	secondNum := rand.Intn( 100 )
	time.Sleep( 2 * time.Second )
	var result = firstNum * secondNum
	fmt.Printf( "result of multiplying %d and %d is: %d \n", firstNum, secondNum, result )
}

