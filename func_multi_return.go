package main

import (
	"fmt"
	"math"
)

// func manyReturns() (int, bool, string) {
// 	return 1, true, "hello"
// }

// func main() {
// 	myInt, myBool, myString := manyReturns()
// 	fmt.Println(myInt, myBool, myString)
// }

func floatParts(number float64) (interger int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)
}
