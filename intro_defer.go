package main

import "fmt"

func intro_defer(){
  fmt.Println("=== Intro Defer ====")
  fmt.Println("// Defer is used to defer execution of a function //")

  printFirst()
  // by add defer keyword the execution of printSecond is deferred until intro_defer / outer func exits
  defer printSecond()
  defer printThird()
  defer printFourth()
}

func printFirst(){
  fmt.Println("First Function")
}

func printSecond(){
  fmt.Println("Second Function")
}

func printThird(){
  fmt.Println("Third Function")
}

func printFourth(){
  fmt.Println("Fourth Function")
}