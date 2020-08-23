package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open( "family.csv" )
	if err != nil {
		log.Fatal( err )  
	}

	scanner := bufio.NewScanner( file )

	for scanner.Scan() {
		var age int
		var name string
		n, err := fmt.Sscanf( scanner.Text(), "%s is %d years old\n", &name, &age )
		if err != nil {
			panic( err )
		}
		if n == 2 {
			fmt.Printf( "Here is name: %s, and here is age: %d \n", name, age )
		}
	}
}

