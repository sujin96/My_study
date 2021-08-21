//1.매개변수에 구조체타입으로 사용하는 예제
//2.반환값을 구조체타입으로 사용하는 예제

package main

import "fmt"

type part struct { //part 타입 지정
	description string
	count       int
}

func showInfo(p part) { //part라는 타입을 가진 매개변수 p 선언
	fmt.Println("Description :", p.description)
	fmt.Println("Count: ", p.count)
}

func minimumOrder(description string) part { //매개변수는 string 타입으로 선언, 반환값을 part 타입으로 선언
	var a part
	a.description = description
	a.count = 100
	return a
}

func main() {
	var bolts part                  //part 타입을 가진 bolts라는 변수 선언
	bolts.description = "Hex bolts" //타입의 필드에 값을 전달
	bolts.count = 24
	showInfo(bolts)
	d := minimumOrder("Hex bolt")
	fmt.Println(d.description, d.count)
}
