package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("x:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i:", i)
	// }

	names := []string{"mario", "luigi", "yoshi", "peach"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("The position at index %v is %v.\n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The value is %v.\n", value)
		// value = "new string" // It doesn't change the actual value.
	}

	fmt.Println(names)
}
