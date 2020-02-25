package main

import ("fmt")

func intro_loop(){
  fmt.Println("=== Intro Loop (For & Range) ===")
  
  simple_for_loop()

  nested_for_loop()

  infinite_loop()

  // see for with range in intro_slice

}

func simple_for_loop(){
  fmt.Println("== Simple Loop ")
  for i := 0; i < 10; i++ {
    if i % 2 == 0 {
      continue
    }
    fmt.Println("Printing odd number only:",i)
  }
}

func nested_for_loop(){
  fmt.Println("== Nested Loop ==")
  for i := 0; i <= 5; i++ {
    for j :=0; j <= 2; j++ {
      fmt.Printf("Outer loop %d, Inner loop %d\n", i, j)
    }
  }
}

func infinite_loop(){
  i := 0
  for {
    if i > 20 {
      break
    }
    fmt.Print("*")
    i++
  }

  fmt.Println("")
}
