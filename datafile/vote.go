//슬라이스를 이용한 개표시스템
//count는 파일에서 각 라인이 나타나는 횟수를 집계합니다.
package main

import (
	"fmt"
	"log"
	"sujinpackage/datafile"
)

func main() {
	lines, err := datafile.GetStrings("vote.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string //후보자들의 이름 목록을 저장할 변수
	var counts []int   // 각 이름이 나타나는 횟수를 저장
	for _, line := range lines {
		matched := false
		for i, name := range names { //names 슬라이스를 순회
			if name == line { //읽어 온 라인이 현재 이름과 일치하면...
				counts[i]++    //카운트를 하나 증가
				matched = true //발견했음을 표시
			}
		}
		if matched == false { //일치하는 이름이 없으면 새로 추가
			names = append(names, line) //새로운 이름을 추가
			counts = append(counts, 1)  //카운트 값도 하나 추가(첫 등장이므로 초깃값으로 1을 사용)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
