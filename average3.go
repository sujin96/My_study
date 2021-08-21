//가번인자함수에 슬라이스 전달하여 평균 계산하기
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 { //가변인자를 매개변수로 선언
	var sum float64 = 0 //인자들의 합계를 저장할 변수를 선언
	for _, number := range numbers {
		sum += number //인자 값을 합계에 더하기
	}
	return sum / float64(len(numbers)) //합계를 인자의 개수로 나눠 평균을 구함
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}
