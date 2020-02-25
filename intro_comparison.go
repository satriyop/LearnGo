package main

import "fmt"

func intro_comparison() {
  fmt.Println("=== Intro Comparison ===")
  // isBool is declare in package level hence can be accessed from here
  fmt.Println("Comparing boolean : ", isBool == true)

	a := 10
	b := 10
	c := 100
	fmt.Println("Comparing integer equality : ", a == b)
	fmt.Println("Comparing integer inequality : ", a != b)
	fmt.Println("Comparing integer greater or more than :", a >= c)
}