package main

import "fmt"

func intro_bitshifting(){
  fmt.Println("=== Intro BitShifting ===")
  a := 1 // 0001 
  b := a << 1 // 0010 = 2
  c := 5 // 0101
  d := c << 1 //1010 = 10
  e := c << 2 // 00010100 = 20
  fmt.Println("Bit shifted : ",b)
  fmt.Println("Bit shifted : ",d)
  fmt.Println("Bit shifted 2 times :",e)
}