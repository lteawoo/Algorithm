package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var countArr []int
	countArr = make([]int, 'Z'-'A'+1)

	var max = 0
	var maxChar byte
	for {
		char, _ := reader.ReadByte()

		//문자열의 끝(10, 13 각각 Line Feed, Return Carriage)을 체크하기 위한.. 그외의 문자가 아닌것도 겸사겸사 체크하여 중지하던지 함
		if (('A' > char) || ('Z' < char)) && (('a' > char) || ('z' < char)) {
			break
		}
		//소문자면 대문자로 변환(문자가 'a' 이상 -> 97 이상 -> 소문자면)
		if char >= 'a' {
			char = char - ('a' - 'A')
		}

		idx := int(char - 'A')

		countArr[idx]++

		//최대 카운트 계산
		if max < countArr[idx] {
			max = countArr[idx]
			maxChar = char
		} else if max == countArr[idx] {
			maxChar = '?'
		}
	}

	fmt.Fprintf(writer, "%c", maxChar)
}
