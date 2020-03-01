package main

import "fmt"

func intro_closure(){
  fmt.Println("=== Intro Closure ===")

  {
    x := 100
    fmt.Println("# This is a closure, enclose the scope of x in this code block only",x)
  }

  a := incrementor()
  b := incrementor()

  fmt.Println("# Closure keep the state of the variable to a certain scope")
  fmt.Println(a())
  fmt.Println(a())
  fmt.Println(a())

  // when b is called it has different scope than a
  fmt.Println(b())
  fmt.Println(b())
  
  fmt.Println(a())
}

func incrementor() func() int {
  x := 0
  return func() int {
    x++
    return x
  }
}