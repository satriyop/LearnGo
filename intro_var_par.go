package main

import "fmt"

func intro_var_par() {
  fmt.Println("=== Variadic Parameter===")
  fmt.Println("Takes infinite number of parameter as slice")


  variadic_int(2, 3, 4, 5, 6)
  variadic_string("hello", "world", "go")
}

func variadic_int(x ...int){
  fmt.Println(x)
  fmt.Printf("%T\n", x)

  sum := 0
  for _, v := range x {
    sum = sum + v
  }
  fmt.Println(sum)
}

func variadic_string(s ...string){
  fmt.Println(s)
  fmt.Printf("%T\n", s)
}