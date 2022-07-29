package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	firstName, secondName := getInitials("Turin Turambar")
	firstName2, secondName2 := getInitials("Bilbo Baggins")
	firstName3, secondName3 := getInitials("Glaurung")

	fmt.Println(firstName, secondName)
	fmt.Println(firstName2, secondName2)
	fmt.Println(firstName3, secondName3)
}
