package main

import "fmt"

func intro_array(){
  fmt.Println("=== Intro Array ===")
  // array declaration
   
  // default zero value for element, in this case 0 since it is an array of integers

  var x [5]int
  fmt.Println(x)

  // set value of array's element
  x[2] = 100
  fmt.Println(x)

  // get value of array's element
  fmt.Println(x[2])

  // get the length of an array
  fmt.Println(len(x))

  // out of bound
  // fmt.Println(x[6])

  // short assignment or Composite Literal
  y := [3]int {1,2,3}
  fmt.Println(y)

}