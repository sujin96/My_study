//average 프로그램은 숫자들의 평균을 계산합니다. (소수의 평균 계산)
package main

import "fmt"

func main() {
	//소수들을 배열로 numbers에 선언
	numbers := [3]float64{71.8, 56.2, 89.5}
	//sum이라는 변수를 선언
	var sum float64
	//인덱스값은 무시하고 소수값만 출력
	for _, number := range numbers {
		//출력한 소수들을 모두 더함
		sum += number
	}
	//배열의 갯수를 count 변수에 선언
	count := float64(len(numbers))
	//더한 소수들을 배열의 개수로 나눠준 값을 출력
	fmt.Printf("avarage : %0.2f\n", sum/(count))
}
