package main

import "fmt"

// Interface is also a type
// Define Interface
type Logger interface {
  logging()
}

// Define type 
type LogToDatabase struct {
  description string
  success bool
  count int
}

type LogToFile struct {
  description string
  success bool
  count int
  fs string
}


func intro_interface(){
  fmt.Println("=== Intro Interface ===")

  logToMySQL := LogToDatabase{
    description: "Saving logs to a MySQL DB",
    success: false,
    count: 0,
  }

  logToFAT32File := LogToFile {
    description: "Saving logs to filesystem",
    success: false,
    count: 0,
    fs: "FAT32",
  }

  // invoke func that its parameter is interface Logger
  startLogging(logToMySQL)

  logSummary(logToMySQL)
  logSummary(logToFAT32File)

}

// create method fort LogToDatabase type
func (logger LogToDatabase) logging() {
  fmt.Println("Writing Logs To DB")
  logger.success = true
  logger.count++
}

func (logger LogToFile) logging() {
  fmt.Println("Writing Logs File")
  logger.success = true
  logger.count++
}

func startLogging(l Logger){
  l.logging()
}
// a function that accept Logger Interface, callable to type LogToDatabase
func logSummary(l Logger){
  switch l.(type) {
    case LogToDatabase:
      fmt.Println("You are logging to a database")
      // Using Assertion
      fmt.Println(l.(LogToDatabase).description)
    case LogToFile:
      fmt.Println("You are logging to a file")
      fmt.Println(l.(LogToFile).description)
  }
}
