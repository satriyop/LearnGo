package main

import "fmt"

func intro_recursion(){
  fmt.Println("=== Intro Recursion ===")

  // most common recursion ex is factorial
  r := factorial_recursion(4)
  fmt.Println(r)

  l := factorial_loop(4)
  fmt.Println(l)

}

func factorial_recursion(n int) int{
  if n == 0 {
    return 1
  } 
  
  return n * factorial_recursion(n - 1)
}

func factorial_loop(n int) int {
  total := 1
  for ; n > 0; n-- {
    total *= n
  }

  return total
}