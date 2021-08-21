//map을 이용한 개표시스템
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
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
