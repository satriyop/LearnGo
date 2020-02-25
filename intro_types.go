package main

import (
  "fmt"
)

var isBool bool

func intro_types() {
  fmt.Println("=== Intro to Types in Go ===")
  
  // BOOLEAN TYPES
	fmt.Println("Zero value boolean is ", isBool)
	isBool = true
	fmt.Println("Changing value in the same type, in this case is bool", isBool)

	// NUMERIC TYPES
	y := 5
	z := 5.5
	maxInt := 9223372036854775807

	fmt.Println(y, z)
	// use print format
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	// using more than max value will capped
	fmt.Println(maxInt+1)

  // STRING TYPES
	s := "s"

  // Conversion not casting
	bs := []byte(s)
	S := "S"
	BS := []byte(S)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%T\n", BS)
	 
	// look at UTF-8 Table to prove
	fmt.Println(bs)
	fmt.Println(BS)
	
}