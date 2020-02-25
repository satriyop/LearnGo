package main

import (
  "fmt"
)

// Using var will expose this variabel in package level
var global_variable = "Package Level Variable"


func intro_variable() {
  fmt.Println("=== Intro to Variable ===")
  
  // function scope Variable
  a := 1
  b := 2
  fmt.Println("This is a variadic variabel", a, b)
}