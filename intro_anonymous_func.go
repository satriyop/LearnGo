package main

import "fmt"

func intro_anonymous_func(){
  fmt.Println("=== Intro Anonymous Function")
  // Anonymous Func that will self exec immediately
  func(){
    fmt.Println("I am Anonymous Func")
  }()

  // Anonymous Func that take arguments
  func(x int){
    fmt.Println("I am Anonymous func with arguments", x)
  }(10)
}