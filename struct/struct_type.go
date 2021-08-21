package main

import "fmt"

type subscriber struct { //subscriber라는 타입을 지정
	name   string
	rate   float64
	active bool
}

func main() {
	var subscriber1 subscriber //해당 타입을 가지는 변수 선언
	var subscriber2 subscriber //해당 타입을 가지는 변수 선언

	subscriber1.name = "Aman singh" //타입의 필드를 가져와서 값을 할당
	subscriber2.name = "Beth Ryan"  //타입의 필드를 가져와서 값을 할당

	fmt.Println("Name :", subscriber1.name)
	fmt.Println("Name :", subscriber2.name)
}
