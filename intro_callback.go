package main

import "fmt"

func intro_callback(){
  fmt.Println("=== Intro Callback ===")
  fmt.Println("# You can pass function as argument to other func #")

  xi := []int {1,2,3,4,5,6,7,8,9}
  sum_result := sum(xi...)
  fmt.Println(sum_result)

  sum_result = sum_even_number(sum, xi...)
  fmt.Println(sum_result)
}

func sum(xi ...int) int {
  total := 0
  for _,v := range xi {
    total += v
  }
  return total
}

func sum_even_number(fn func(xi ...int) int, xi ...int) int {
  total := 0
  for _, v := range xi {
    if v % 2 == 0 {
      total += v
    }
  }
  return total
}