//플레이어가 난수를 맞추는 게임입니다.

//조건
/*플레이어가 추측할 1에서 100 사이의 난수를 목푯값으로 지정합니다.
플레이어가 추측한 숫자를 입력할 수 있도록 프롬프트를 띄우고 입력받은 추측 값을 저장합니다.
플레이어가 추측한 숫자가 목푯값보다 낮으면 "Oops.Your guess was LOW."를, 높으면 "Oops.Your guess was HIGh."를 출력합니다.
플레이어는 최대 10번까지 추측할 수 있습니다. 추측을 할 때마다 플레이어에게 남은 횟수를 알려줍니다.
플레이어가 추측한 숫자가 목푯값과 같으면 "Good job! You guessed it!"을 출력하고 프롬프트를 종료합니다.
최대 추측 횟수 안에 목푯값을 맞히지 못하면 "Sorry. You didn't guess my number. It was: [targer]"을 출력합니다.*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() //현재 날짜 및 시간을 정숫값으로 가져옵니다.
	rand.Seed(seconds)           //난수 생성기를 시드 합니다.
	target := rand.Intn(100) + 1 //1~100까지 난수를 생성합니다.
	fmt.Println("제가 1~100중에 랜덤으로 숫자하나를 뽑았습니다.")
	fmt.Println("맞춰보세요")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) //키보드 입력을 읽기위한 bufio.Reader를 생성합니다.

	success := false //루프가 시작되기 전에 "success"를 선언 하면 루프가 종료된 후에도 스코프를 유지할 수 있습니다.
	for i := 0; i < 10; i++ {
		fmt.Println("기회가", 10-i, "번 남았습니다") //10에서 추측 횟수를 차감한 남은 추측 횟수를 플레이어에게 알려 줍니다.

		fmt.Print("추측하는 값: ")                 //추측한 값을 물어봅니다.
		input, err := reader.ReadString('\n') //사용자가 엔터키를 누를때까지 입력한 내용을 읽어 옵니다.
		if err != nil {
			log.Fatal(err) //에러가 발생하면 에러 메시지를 출력하고 프로그램을 종료합니다.
		}
		input = strings.TrimSpace(input)  //줄 바꿈 문자를 제거합니다.
		guess, err := strconv.Atoi(input) //입력 문자열을 정숫값으로 변환합니다.
		if err != nil {
			log.Fatal(err) //에러가 발생하면 에러 메시지를 출력하고 프로그램을 종료합니다.
		}

		if guess < target {
			fmt.Println("정답보다 낮게 추측했어요")
		} else if guess > target {
			fmt.Println("정답보다 높게 추측했어요")
		} else {
			success = true
			fmt.Println("맞췄습니다!! 정답은", target, "이었습니다.")
			break
		}

	}
	if !success {
		fmt.Println("모든 기회를 소진하였습니다. 답은", target, "이었습니다.")
	}
}
