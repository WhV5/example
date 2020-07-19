package main

import "fmt"

func main() {
	m1 := map[string]int{
		"Henry": 12,
		"Bob":   18,
		"Nancy": 22,
	}
	fmt.Printf("%v\n", m1)

	m2 := make(map[string]int)
	m2["Henry"] = 12
	m2["Bob"] = 18
	m2["Nancy"] = 22
	fmt.Printf("%v\n", m2)

	var m3 map[string]int
	if m3 == nil {
		fmt.Println("m3 is null")
	}

	m4 := make(map[string]int)
	if m4 == nil {
		fmt.Println("m4 is null")
	}
}
