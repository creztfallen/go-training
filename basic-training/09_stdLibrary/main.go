// Standard Library packages: https://golang.org/pkg/
package main

import (
	"fmt"
	//strings
	"sort"
)

func main() {
	// greeting := "Hello there people."

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// replaceAll doesn't change the original string
	// fmt.Println("Original value =", greeting)

	ages := []int{54, 22, 24, 30, 75, 60, 51, 25, 7}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Geralt", "Triss", "Jaskier", "Plotka", "Dijkstra"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Dikstra"))
}
