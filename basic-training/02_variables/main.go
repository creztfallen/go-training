package main

import "fmt"

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

//The := constructor is never available outside of functions.

// Constants

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

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

func main() {
	const Greeting = "ハローワールド"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
}
