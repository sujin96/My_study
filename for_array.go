//배열 순회

package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	//for문을 이용하여 배열을 순회하면서 index값과 문자열을 출력
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	//for range문을 이용한 순회 - 배열의 인덱스값을 알아야 하고 실수할 여지가 있기때문에 안전한 방법인 for range문을 사용한다.
	//index = 인덱스 값
	//note = 그 인덱스의 문자열

	//배열을 0~6번 인덱스까지 번호와 문자열을 순회하면서 출력
	for index, note := range notes {
		fmt.Println(index, note)
	}

	//인덱스만 출력
	for index, _ := range notes {
		fmt.Println(index)
	}

	//문자열만 출력
	for _, note := range notes {
		fmt.Println(note)
	}
}
