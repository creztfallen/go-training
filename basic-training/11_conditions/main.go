package main

import "fmt"

func main() {
	// age := 25

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 40")
	// }

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at position", index)
			continue // reiterate the loop at this point
		}

		if index > 2 {
			fmt.Println("Breaking at position", index)
			break // break the loop at this point
		}

		fmt.Printf("The value at position %v is %v. \n", index, value)
	}
}
