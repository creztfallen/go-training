package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning, %s! \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye, %s! \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	diameter := r * r
	return math.Pi * (diameter)
}

func main() {
	// sayGreeting("Geralt")
	// sayBye("Jaskier")

	cycleNames([]string{"Bifur", "Bofur", "Bombur"}, sayGreeting)
	cycleNames([]string{"Bifur", "Bofur", "Bombur"}, sayBye)

	a1 := circleArea(9.7)
	a2 := circleArea(12)

	// fmt.Println(a1, a2)
	fmt.Printf("a1:%0.2f\n", a1)
	fmt.Printf("a2:%0.2f\n", a2)
}

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
