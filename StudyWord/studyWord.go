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
	fmt.Fprintf(writer, "%d\n", len(word))

	for i := 0; i < len(word); i++ {
		var char = byte(word[i])

		if char >= 97 {
			char -= 32
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

	fmt.Fprintf(writer, "%c %d", maxKey, max)
}
