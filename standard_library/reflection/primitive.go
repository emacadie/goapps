package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	ourTitle := "the go standard library"
	newTitle := properTitle( ourTitle )
	fmt.Println( "our function properTitle changed '", ourTitle, "' to: '", newTitle, "'" )
	fmt.Println( "Calling doubleOurNumber( 3 ): ", doubleOurNumber( 3 ) )
	fmt.Println()
	timedTitleFunc := MakeTimedFunction( properTitle ).( func( string ) string )
	timedTitle := timedTitleFunc( ourTitle )
	fmt.Println( "timedTitle is: '" + timedTitle + "'" )
	timedMultFunc := MakeTimedFunction( doubleOurNumber ).( func( int ) int )
	timedMultVar := timedMultFunc( 4 )
	fmt.Println( "timedMultFunc( 4 ) is: ", timedMultVar )
}

func properTitle( input string ) string {
	// from https://golangcookbook.com/chapters/strings/title
	words := strings.Fields( input )
	smallwords := " a an on the to "
	
	for index, word := range words {
		if strings.Contains( smallwords, " "+word+" " ) {
			words[ index ] = word
		} else {
			words[ index ] = strings.Title( word )
		}
	}
	return strings.Join( words, " " )
} // properTitle

func doubleOurNumber( argNum int ) int {
	time.Sleep( 1 * time.Second )
	return argNum * 2
}

func MakeTimedFunction( f interface{} ) interface{} {
	rf := reflect.TypeOf( f )
	if rf.Kind() != reflect.Func {
		panic( "Expecting a function" )
	}
	vf := reflect.ValueOf( f )
	wrapperF := reflect.MakeFunc( rf, func( in []reflect.Value ) []reflect.Value {
		start := time.Now()
		out   := vf.Call( in )
		end   := time.Now()
		fmt.Printf( "Calling %s took %v\n", runtime.FuncForPC( vf.Pointer() ).Name(), end.Sub( start ) )
		return out
		} )
	return wrapperF.Interface()
}
	


