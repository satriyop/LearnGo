package main

import "fmt"

func intro_func(){
  fmt.Println("=== Intro Function ===")
  fmt.Println("func(r receiver) identifier(parameter(s)) (return(s)) {...}")
  printName("Satriyo")

  s1 := returnLastName("Pamungkas")
  fmt.Println(s1)

  a, b := multipleReturn("BackEnd")
  fmt.Println(a, b)

}

// Everything in Go is PASSED BY VALUE
func printName(s string) {
  fmt.Println("Hello", s)
}

// Return from Func
func returnLastName(s string) string {
  return s
}

func multipleReturn(s string) (string, bool) {
  if s == "BackEnd" {
    return s, true
  }
  return "Sorry", false
}
