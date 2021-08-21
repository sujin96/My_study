//타입을 지정하고 매개변수와 반환값을 둘다 사용하는 예제
package main

import "fmt"

type subscriber struct { //타입 지정
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) { //지정한 타입의 매개변수 선언
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber { //지정한 타입을 반환값으로 사용
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman singh") //subscriber1 변수는 구조체가 아닌 구조체 포인터 ->  리턴을 포인터로 받았기 때문
	applyDiscount(subscriber1)                     //이미 포인터 변수이므로 주소 연산자는 사용하지 않음
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}

/* main 함수에서는 defaultSubscriber에 구독자 이름을 전달해 새로운 subscriber 구조체 값을 생성하고 있습니다.
구독료를 할인받은 구독자는 rate 필드를 applyDiscount에서 설정해 줍니다. 값이 채워진 subscriber 구조체를 printInfo 함수로 전달하면 구조체에
저장된 값들을 출력할 수 있습니다. */
