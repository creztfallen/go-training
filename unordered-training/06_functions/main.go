package main

import "fmt"

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

func place(city string) (country, continent string) {

	switch city {
	case "Rio de Janeiro", "Copacabana", "Lapa":
		country, continent = "Brasil", "South America"
	case "Tokyo", "Okinawa":
		country, continent = "Japan", "Asia"
	case "Yalova", "Aydin", "Istanbul":
		country, continent = "Türkiye", "Asia"
	default:
		country, continent = "Unknown", "Unknown"
	}
	return
}

func main() {
	country, continent := place("Aydin")
	fmt.Printf("Mateus %s'de, %s'da yasiyor", country, continent)
}
