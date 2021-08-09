package main

import "fmt"

/*함수 사용 전*/
// func main() {
// 	var width, height, area float64
// 	width = 4.2
// 	height = 3.0
// 	area = width * height
// 	fmt.Printf("%.2f liters needed\n", area/10.0)

// 	width = 5.2
// 	height = 3.5
// 	area = width * height
// 	fmt.Printf("%.2f liters needed\n", area/10.0)
// }

/*함수 사용 후*/
func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(4.2, 3.5)
	paintNeeded(5.0, 3.3)
}
