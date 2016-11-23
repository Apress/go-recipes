package main

import (
	"flag"
	"fmt"
)

func main() {

	strPtr := flag.String("name", "Shiju", "a string")
	numbPtr := flag.Int("num", 25, "an int")
	boolPtr := flag.Bool("enable", false, "a bool")
	var num int
	flag.IntVar(&num, "num", 30, "an int")	
	// Parse parses flag definitions from the argument list.
	flag.Parse()
	// Get the values for pointers
	fmt.Println("name:", *strPtr)
	fmt.Println("num:", *numbPtr)
	fmt.Println("enable:", *boolPtr)
	// Get the value from a variable
	fmt.Println("num:", num)
	// Args returns the non-flag command-line arguments.
	fmt.Println("arguments:", flag.Args())
}
