package main

import "fmt"

func intro_return_a_func(){
  fmt.Println("=== Intro Returning a Function ===")

  f1 := f_generator()
  f1()

  f2 := f_int_generator()
  fmt.Printf("The returning func will return type of %T\n", f2)

  f2int := f2()
  fmt.Printf("The returning func exec'ed and return type of %T\n", f2int)

  // or you can do this if you dont want to assign to variable
  fmt.Println(f_int_generator()())
  
}

func f_generator() func() {
  fn := func(){
    fmt.Println("First Func Return")
  }
  return fn
}

func f_int_generator() func() int {
  fn := func() int {
    fmt.Println("Second Func that will return int")
    return 1984
  }

  return fn
} 