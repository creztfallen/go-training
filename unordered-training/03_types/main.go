package main

import (
	"fmt"
	"time"
)

/*
bool
string

Numeric types:

uint        either 32 or 64 bits
int         same size as uint
uintptr     an unsigned integer large enough to store the uninterpreted bits of
            a pointer value
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers
            (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32 (represents a Unicode code point)
*/

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
