package main

import (
	"fmt"
)

type Person struct {
  firstName string
  lastName string
  age int
}
// Embedding struct Person to Employee
type Employee struct {
  Person
  isActive bool
  salary float64
}


type Hero struct {
	RealName string
	Health   int
}

func intro_structures() {
	fmt.Println("=== Intro Structures ===")

  fmt.Println("== Create a value for a struct")
  p1 := Person {
    firstName: "Satriyo",
    lastName: "Pamungkas",
    age: 45,
  }

  fmt.Println(p1.firstName, p1.lastName, p1.age)

  

	// This is valid
	superman := Hero{}
	fmt.Println(superman)

	// Should be oneliner to be valid
	ironman := Hero{RealName: "Elon Musk"}
	fmt.Println(ironman)

	// Don't forget comma trailing
	spiderman := Hero{
		RealName: "Peter Parker",
		Health:   110,
	}
	fmt.Println(spiderman)

	// Or can be simple oneliner
	batman := Hero{RealName: "Bruce Wayne", Health: 80}
	fmt.Println(batman)

	// Also, without the field name, remember apply this on few fields for clarity
	thor := Hero {"Chris Hemworth", 10}
	fmt.Println(thor)





// Embedded Struct
  fmt.Println("== Embedded Struct ==")
  e1 := Employee {
    Person: p1,
    isActive: true,
    salary: 1000.5,
  }

  fmt.Println(e1)
  // or using composite literal to create embedded struct
  e2 := Employee {
    Person: Person {
      firstName : "Bumi",
      lastName: "Pamungkas",
      age: 7,
    },
    isActive: false,
    salary: 10,
  }

  fmt.Println(e2)



  // using Method in Struct with pointer
  e1.layoff()
  fmt.Println("After execute method layoff", e1)


  // Anonymous Struct
  fmt.Println("== Anonymous Struct ==")

  d := struct{
    role string
    skill []string
  } {
    role: "Back End Developer",
    skill: []string {"php", "java", "golang",},
  }

  fmt.Println(d)

	
  // === Pointer in struct ===//
	
	// Using Pointer and Address to Operator & Pointer as Argument to Heal Function
  fmt.Println("== Pointer Struct ==")
	wonder_woman := &Hero{"Galgadot", 9000}
	fmt.Println(wonder_woman)

	// Apply Healing to beloved Wonder Woman
	Heal(wonder_woman)
	fmt.Println("Wonder woman health now is :",wonder_woman.Health)

}

// === function with struct === //

// Go passes arguments to a function: as copies. hence need pointer to modify the real value

func Heal(h *Hero) {
  fmt.Println("== Function with struct as argument ==")
	h.Health += 1000
}

// === Method with struct ===//
// Struct will be the receiver of a function
func (e *Employee) layoff(){
  e.isActive = false;
  fmt.Println("== Method in struct ==")
}
