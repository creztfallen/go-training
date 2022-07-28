package main

import (
	"fmt"
	"time"
)

// var (
// 	goIsFun bool       = true
// 	maxInt  uint64     = 1<<64 - 1
// 	complex complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	const f = "%T(%v)\n" // type conversion where t is the type and v is the value being converted to the type.
// 	fmt.Printf(f, goIsFun, goIsFun)
// 	fmt.Printf(f, maxInt, maxInt)
// 	fmt.Printf(f, complex, complex)
// }

// Type Conversion

// i := 42
// f := float64(i)
// u := uint(f)

// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)

// func main() {
// 	fmt.Printf("%d, %f, %d", i, f, u)
// }

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

type Point struct {
	X, Y int
}

func main() {

	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
	}

	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)", event.Date, event.Lat, event.Lon)
	fmt.Println(Bootcamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	})

	p := Point{1, 2}  //has type Point
	q := &Point{1, 2} // has type *Point
	r := Point{X: 1}  // Y:0 is implicit
	s := Point{}      // X:0 and Y:0

	fmt.Println(p, q, r, s)
}
