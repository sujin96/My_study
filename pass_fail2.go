package main

import "fmt"

func main() {
	var a float32
	fmt.Print("Enter a grade: ")
	fmt.Scan(&a)

	var status string
	if a >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", a, "is", status)
}
