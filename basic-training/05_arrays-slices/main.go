package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{22, 24, 12}
	ages := [3]int{22, 24, 12}

	names := [4]string{"Dark Souls", "Skyrim", "Dragon's Dogma", "The Witcher"}
	// names[1] = "Shovel Knight"

	fmt.Println(ages, len(ages)) // len = length
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	scores := []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85) //append() alone just returns a new arr but doesn't change the original

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]  //ranges are inclusive for the beginning but exclusive for the end
	rangeTwo := names[1:]   //when omitted, the range limit is set to the get from the set up starter until the last array element.
	rangeThree := names[:3] // goes from the first element until the last -1 (exclusive)
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Deep Down")
	fmt.Println(rangeOne)
}
