package main

import "fmt"

func intro_channel(){
    fmt.Println("=== Intro Channel ===")
    // Make channel
    message := make(chan string)
    
    go func (){
        // Send Value to channel
        message <- "Ping Pong"
        message <- "Pam Pam"
    }()

    // Receive value from channel
    msg := <-message
    msg2 := <- message

    fmt.Println(msg)
    fmt.Println(msg2)

    channel_buffering()
}

func channel_buffering(){
    message := make(chan string, 2)

    message <- "buffered"
    message <- "channel"

    // Without concurent receiver
    fmt.Println(<-message)
    fmt.Println(<-message)
}

