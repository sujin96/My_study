//외부구조체에 필드타입으로 내부구조체 사용
//익명 구조체 필드 사용
//임베딩
package main

import "fmt"

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address //필드 이름 없이 익명 구조체 필드로 사용할 수 있음
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address //Address타입을 필드 타입으로 사용
}

type Address struct { //Address 라는 사용자지정 타입 만들기
	Street     string
	City       string
	State      string
	PostalCode string
}

func main() {
	address := Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}

	var address2 Subscriber //익명 구조체 필드는 내부구조체가 외부구조체로 임베딩되어 마치 외부 구조체에 속해있는 것처럼 사용가능
	address2.City = "seoul"
	fmt.Println(address2.City)

	subscriber := Subscriber{Name: "Aman Singh"}
	subscriber.Address = address
	fmt.Println(subscriber.Name)
	fmt.Println(subscriber.Address)

}
