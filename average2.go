//가번인자함수를 사용해 평균 계산하기
package main

import "fmt"

func average(numbers ...float64) float64 { //가변인자를 매개변수로 선언
	var sum float64 = 0 //인자들의 합계를 저장할 변수를 선언
	for _, number := range numbers {
		sum += number //인자 값을 합계에 더하기
	}
	return sum / float64(len(numbers)) //합계를 인자의 개수로 나눠 평균을 구함
}

func main() {
	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))
}
