package main

import "fmt"

func intro_map() {
  fmt.Println("=== Intro Slice ===")

	m := map[string]int{
		"Satriyo": 35,
		"Bumi":    5,
		"Gundala": 100,
	}

  // Accessing map via its key
	fmt.Println(m["Bumi"])

  // Loop over the map
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Value not exist")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

  // Add to map
  m["Banyu"] = 8
  

  // Delete elemen in the map
	delete(m, "Gundala")
	for k, v := range m {
		fmt.Println(k, v)
	}

  // Safe delete element with check if it exist
  if _, ok := m["Banyu"]; ok {
    fmt.Println("Deleting ")
    delete(m, "Banyu")
  }
}