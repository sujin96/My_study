//txt 파일을 불러와서 print로 찍는 코드
//txt파일은 같은 디렉토리에 위치해야함
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") //읽을 데이터 파일을 연다.

	//파일을 여는 도중에 에러가 발생하면 에러를 보고하고 프로그램을 종료
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file) //파일에 대한 스캐너를 생성

	//파일 끝에 도달해 scanner.Scan이 false를 반환할때까지 순회한다.
	for scanner.Scan() { //파일에서 한 줄을 읽음
		fmt.Println(scanner.Text()) //한줄의 내용을 출력
	}

	err = file.Close() //파일을 닫아 자원을 반환합니다.

	//파일을 닫는 도중에 에러가 발생하면 에러를 보고하고 프로그램을 종료
	if err != nil {
		log.Fatal(err)
	}

	//파일의 내용을 읽는 도중에 에러가 발생하면 에러를 보고하고 프로그램을 종료
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
