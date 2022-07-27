package main

import (
	"fmt"
)

//VARIABLES

// var (
// 	name string
// 	age int
// 	location string
// )

// var (
// 	name, location string
// 	age int
// )

// var name string
// var age int
// var location string

// var (
// 	name string = "Seppä"
// 	age int = 57
// 	location string = "Finland"
// )

//Inferred Typing
// var (
// 	name = "Seppä"
// 	age = 57
// 	location = "Finland"
// )

// var (
// 	name, location, age = "Seppä", "Finland", 57
// )

//Short Assignment Statement (work inside functions)
// func main() {
// 	name, location := "Seppä", "Finland"
// 	age := 57
// 	fmt.Printf("%s (%d) of %s", name, age, location)
// }

//Functions inside variables
// func main() {
// 	action := func() {
// 		fmt.Println("Opa, bão?")
// 	}
// 	action()
// }

// The := constructor is never available outside of functions.

//CONSTANTS

// const Pi = 3.14
// const (
// 	StatusOK = 200
// 	StatusCreated = 201
// 	StatusAccepted = 202
// 	StatusNonAuthoritativeInfo = 203
// 	StatusNoContent = 204
// 	StatusResetContent = 205
// 	StatusPartialContent = 206
// )

// const (
// 	Pi    = 3.14
// 	Truth = false
// 	Big   = 1 << 62
// 	Small = Big >> 61
// )

// func main() {
// 	const Greeting = "ハローワールド"
// 	fmt.Println(Greeting)
// 	fmt.Println(Pi)
// 	fmt.Println(Truth)
// 	fmt.Println(Big)
// }

//PRINTING CONSTS AND VARS

// func main() {
// 	age := 24
// 	name := "Mateus"
// 	/*fmt.Println prints the passed in variables’
// 	values and appends a newline.*/
// 	fmt.Println(age)
// 	fmt.Printf("%s is %d years old.", name, age)
// }

//FUNCTIONS, SIGNATURE, RETURN VALUES, NAMED RESULTS

// func add(x int, y int) int {
// 	return x << y
// }

// func add(x, y int) int {
// 	return x << y
// }

// func main() {
// 	fmt.Println(add(2, 9))
// }

// func place(city string) (string, string) {
// 	var (
// 		country, continent string
// 	)

// 	switch city {
// 	case "Rio de Janeiro", "Copacabana", "Lapa":
// 		country, continent = "Brasil", "South America"
// 	case "Tokyo", "Okinawa":
// 		country, continent = "Japan", "Asia"
// 	case "Yalova", "Aydin", "Istanbul":
// 		country, continent = "Türkiye", "Asia"
// 	default:
// 		country, continent = "Unknown", "Unknown"
// 	}
// 	return country, continent
// }

// func place(city string) (country, continent string) {

// 	switch city {
// 	case "Rio de Janeiro", "Copacabana", "Lapa":
// 		country, continent = "Brasil", "South America"
// 	case "Tokyo", "Okinawa":
// 		country, continent = "Japan", "Asia"
// 	case "Yalova", "Aydin", "Istanbul":
// 		country, continent = "Türkiye", "Asia"
// 	default:
// 		country, continent = "Unknown", "Unknown"
// 	}
// 	return
// }

// func main() {
// 	country, continent := place("Aydin")
// 	fmt.Printf("Mateus %s'de, %s'da yasiyor", country, continent)
// }

//POINTERS

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
	me := &Artist{Name: "Beyza", Genre: "Folk", Songs: 19}
	fmt.Printf("%s released her %dth song.\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs.\n", me.Name, me.Songs)

}
