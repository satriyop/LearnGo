package main

import "fmt"

func intro_func_expression(){
  fmt.Println("=== Intro Func Expression ===")
  // Saving function to a variable (func is 1st class citizen )
  // Func is a Type
  f1 := func(){
     fmt.Println("This is first func expression")
   }
  // invoke the function
  f1()

  f2 := func(x int){
     fmt.Println("This is first func expression with arguments", x)
   }
  // invoke the function
  f2(5)
}