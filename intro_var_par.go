package main

import "fmt"

func intro_var_par() {
  fmt.Println("=== Variadic Parameter===")
  fmt.Println("Takes infinite number of parameter as slice or even zero values")


  variadic_int(2, 3, 4, 5, 6)
  variadic_string("hello", "world", "go")
  variadic_combine("golang", 1, 2, 3)

  // spread the slice as argument
  names := []string {"Sat", "Bum", "Ban"}
  variadic_spreadParameter("hello", names...)
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

// ...T parameter should be in the last order of parameters 
func variadic_combine(s string, x ...int) {
  fmt.Println(s, x)
}

func variadic_spreadParameter(greet string, names ...string){
  fmt.Println(greet, names)
}