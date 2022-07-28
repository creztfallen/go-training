package main

import "fmt"

// Methods are often defined on pointers and not values (although they can be defined on both)
// So you will often store a pointer in a variable as in the example below:

// func main() {
// 	client := &http.Client{}
// 	res, err := client.Get("http://google.com")
// }

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Beyza", Genre: "lo-fi", Songs: 19}
	fmt.Printf("%s released her %dth song.\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs.\n", me.Name, me.Songs)

}
