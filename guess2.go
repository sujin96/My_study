package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix() //현재 날짜 및 시간을 정숫값으로 가져옵니다.
	rand.Seed(seconds)           //난수 생성기를 시드 합니다.
	b := rand.Intn(100) + 1      //1~100까지 난수를 생성합니다.
	fmt.Println("제가 1~100중에 랜덤으로 숫자하나를 뽑았습니다.")
	fmt.Println("맞춰보세요")

	success := false
	for i := 0; i < 10; i++ {
		fmt.Println("기회가", 10-i, "번 남았습니다.")
		var a int
		fmt.Print("추측하는 값:")
		fmt.Scan(&a)

		if a < b {
			fmt.Println("정답보다 낮게 추측했어요")
		} else if a > b {
			fmt.Println("정답보다 높게 추측했어요")
		} else {
			success = true
			fmt.Println("맞췄습니다!! 정답은", b, "이었습니다.")
			break
		}
	}
	if !success {
		fmt.Println("모든 기회를 소진하였습니다. 답은", b, "이었습니다.")
	}
}
