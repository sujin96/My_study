package main

import "fmt"

var subscriber1 struct {
	name   string
	rate   float64
	active bool
}

var subscriber2 struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber1.name = "Aman singh"
	fmt.Println("Name :", subscriber1.name)
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name :", subscriber2.name)
}

//구조체를 사용하면 타입이 다른 여러 값들을 한번에 저장할 수 있음!
//단순히 var ... struct {} 구조로 구조체를 만들 시 같은 코딩을 여러번 해야함
//var 대신 type을 사용해서 새로운 type으로 지정해주는 것이 좋음
