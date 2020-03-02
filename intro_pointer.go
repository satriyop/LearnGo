package main

import "fmt"

func intro_pointer(){
  fmt.Println("=== Intro Pointer ===")
  a := 100
  fmt.Println("The value of variable a is :",a)
  fmt.Println("The memory address for a is :",&a)
  // Type of a with/out the & symbol
  fmt.Printf("The type of a is : %T\n",a)
  fmt.Printf("The type of &a is : %T\n", &a)
  // *int means pointer to an int type
  fmt.Println("int and *int is different type")

  
  b := &a
  fmt.Println("The value of the memory address is ",*b)
  fmt.Printf("The type of b is %T\n", b)

  var c *int = &a
  fmt.Println("This assignment will work too, c memory address is ", c)


  // if we change c, since it is a pointer it will change a 
  *c = 200
  fmt.Println(a, "The value is now changed from ", *b, "and", *c)

  d := *&a
  fmt.Printf("the value of variable d is %v, with type %T\n", d, d)


  // When to use Pointer
  // 1. When you have a large data and dont want to pass by value (copy value eg pass to argument of a func) - go pass everything by value by default.
  // 2. Want to change a value in a memory address.

  z := 1000
  fmt.Println("The value of z is:",z)
  byValue(z)
  fmt.Println("The value of z is not changed because the argument passed to byvalue func is the copied value / pass by value. z is :",z)

  byPointer(&z)
  fmt.Println("The value of z is now changed", z)
}

func byValue(x int){
  fmt.Println("The copied value passed to argument is :",x)
  x = 88
  fmt.Println("The value of argument is changed but not the original z",x)
}

func byPointer(x *int){
  fmt.Println("This is the original value of z", x)
  *x = 99
  fmt.Println("The value of z is now changed to ",x)
}