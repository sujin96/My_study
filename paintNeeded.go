package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 { //너비 값이 유효하지 않은 경우 0과 에러를 반환합니다.
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 { //높이 값이 유효하지 않은 경우 0과 에러를 반환합니다.
		return 0, fmt.Errorf("a height of % 0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil //페인트의 양과 함께 아무 에러도 없음을 나타내는 "nil" 값을 반환합니다.
}

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters neede\n", amount)
}
