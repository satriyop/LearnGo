package main

import "fmt"

const a = 10

// Typed Constant
const (
  b = 100
  c = "Golang"
)

// Untyped Constant
const (
  x int = 1000
  y float64 = 10.10
  z string = "Go Go Go"
)

func intro_constant(){
  fmt.Println("=== Intro to Constant ===")
  fmt.Println("Constant a value :", a)
  fmt.Println("Another way to declare constant :", b,c)
  fmt.Println("You can specify the types when declaring constant :", x, y, z)
}