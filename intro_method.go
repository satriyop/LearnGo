package main

import "fmt"

type user struct {
  id int
  name string
  isActive bool
  sentEmail int
  receivedEmail int
}

func intro_method(){
  fmt.Println("=== Intro Method ===")
  satriyo := user{
    id:1,
    name: "Satriyo",
    isActive: false,
    sentEmail: 0,
    receivedEmail:0,
  }

  kotaro := user{
    id:2,
    name: "Kotaro Minami",
    isActive: false,
    sentEmail: 0,
    receivedEmail:0,
  }

  activateUser(&satriyo)
  fmt.Println("Whats the impact ?",satriyo)

  // calling test method
  kotaro.test()
  // calling sendEmail method
  if satriyo.sendEmail(kotaro) {
    fmt.Println("Email sent successfully")
  }

}

func activateUser(u *user){
  fmt.Println("Before activate", u)
  if u.isActive == false {
    u.isActive = true
    fmt.Println("After activate", u)
  } else {
    fmt.Println("Nothing to activate")
  }
  
  
}

// Attaching function on type user (hence become method)
func (u user) test(){
  fmt.Println(u)
  fmt.Println("Testing Method Done")
}

func (u user) sendEmail(r user) bool {
  r.receivedEmail++
  u.sentEmail++
  fmt.Println("Sent Email Counts",u.sentEmail)
  fmt.Print("Received Email Counts",r.receivedEmail)
  return true
}

