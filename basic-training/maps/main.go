package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		1: "mario",
		2: "luigi",
		3: "toad",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2])

	phonebook[2] = "peach"
	fmt.Println(phonebook[2])

}
