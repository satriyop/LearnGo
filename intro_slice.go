package main

import "fmt"

func intro_slice(){
  fmt.Println("=== Intro Slice ===")
  // Composite Literal
  x := []int {2, 3, 10, 100, 4}
  fmt.Println(x)
  fmt.Println(x[3])

  // slice capacity is depending on the array capacity as the underlying data structure
  x = append(x, 5)
  fmt.Println(x)

  for i,v := range x { 
    fmt.Println(i, v)
  }

  for _, v := range x {
    fmt.Println(v)
  }
  
  
  fmt.Println("== Unpack Slice ===")
  spreadSlice(x...)

}

func spreadSlice(i ...int){
  for i,v := range i {
    fmt.Println(i,v)
  }
}