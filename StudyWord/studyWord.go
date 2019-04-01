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

	var word string
	var retMap map[byte]int
	retMap = make(map[byte]int)
	word, _ = reader.ReadString('\n')

	//10 : Line Feed, 13 : Cariage Return
	//13 10 으로 구성되어 있네..
	for i := 0; i < len(word); i++ {
		char := byte(word[i])

		if char == 10 || char == 13 {
			break
		}

		if char >= 'a' {
			char -= 'a' - 'A'
		}

		_, exists := retMap[char]

		if !exists {
			retMap[char] = 1
		} else {
			retMap[char]++
		}
	}

	var max = 0
	var maxKey byte
	for key, val := range retMap {
		if max < val {
			max = val
			maxKey = key
		} else if max == val {
			maxKey = '?'
		}
	}

	fmt.Fprintf(writer, "%c", maxKey)
}
