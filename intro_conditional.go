package main

import "fmt"

func intro_conditional(){
  fmt.Println("=== Intro Conditional ===")
  intro_if_else()
  intro_switch()
}

func intro_if_else(){
  fmt.Println("== Conditional using if else")
  // simple if
  a := 10
  if a != 10 {
    fmt.Printf("%d is not 10 \n", a)
  }

  // if with var init
  if name:= "Go"; name == "Go" {
    i := 0
    for i < 5 {
      fmt.Print(name + " ")
      i++
    }
    fmt.Println("")
  }

  // if else
  if a <= 3 {
    fmt.Println("a is less than 3")
  } else if a > 3 && a <= 6 {
    fmt.Println("a is between 3 and 6")
  } else {
    fmt.Println("i dont know a")
  }


}

func intro_switch(){
  fmt.Println("== Conditional using switch")
  // Simple
  switch "true"{
    case "falsy":
      fmt.Println("not print")
    case "true":
      fmt.Println("print true")
    case "falsy again":
      fmt.Println("false again")
  }


  // Falthrough
  switch {
    case false:
      fmt.Println("false 1")
    case true:
      fmt.Println("true 1")
      fallthrough
    case false:
      fmt.Println("false 2")
      fallthrough
    case true:
      fmt.Println("true 2")
    case false:
      fmt.Println("false 3")
          case true:
      fmt.Println("true 3")
    case false:
      fmt.Println("false 4")
  }

  // Default value
  switch {
    case false:
      fmt.Println("not printing false value")
    case false:
      fmt.Println("not print false value")
    case false:
      fmt.Println("false again")
    default:
      fmt.Println("default value")
  }

}