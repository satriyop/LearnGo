package main

import "fmt"

func intro_slice(){
  fmt.Println("=== Intro Slice ===")
  // Composite Literal
  x := []int {2, 3, 10, 100, 4}
  fmt.Println(x)
  fmt.Println(x[3])

  x = append(x, 5)
  fmt.Println(x)

  for i,v := range x { 
    fmt.Println(i, v)
  }

  for _, v := range x {
    fmt.Println(v)
  }

}