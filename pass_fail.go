//pass_fail

// 학생이 자신의 백분율 성적을 입력하면 합격 여부를 알려주는 프로그램을 작성
// 합격/불합격 판단 공식 : 성적 60%이상이면 합격, 60% 미만이면 불합격

package main

import ( //"main"함수에서 사용할 패키지를 가져옵니다.
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" //ParseFloat를 사용하기 위해 "strconv" 패키지를 가져옵니다.
	"strings" //TrimSpace를 사용하기 위해 "strings"패키지를 가져옵니다.
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)   // 키보드로부터 텍스트를 읽어오기 위한 "버퍼리더"를 만듭니다.
	input, err := reader.ReadString('\n') //엔터키가 눌리기 직전까지 입력된 모든 문자를 반환합니다.
	if err != nil {
		log.Fatal(err)
	} // "err" 가 nil이 아니면 에러를 보고하고 프로그램을 종료합니다.

	input = strings.TrimSpace(input)            //입력 문자열에서 줄 바꿈 문자를 제거합니다.
	grade, err := strconv.ParseFloat(input, 64) //문자열을 float64 값으로 변환합니다.
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
