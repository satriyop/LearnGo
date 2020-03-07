package main

import (
    "fmt"
    "time"
)

func intro_goroutine(){
    fmt.Println("=== Intro GoRoutine ===")
    for i := 0; i < 5; i++ {
        f_direct()
        go f_goroutine_1()
        go f_goroutine_2()
        fmt.Println(time.Second)
    }
}

func f_direct(){
    fmt.Println("Direct")
}

func f_goroutine_1(){
    fmt.Println("GoRoutine 1")
}

func f_goroutine_2(){
    fmt.Println("GoRoutine 2")
}

