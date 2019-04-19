package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, resultCnt int
	var arr []int
	var char, preChar byte

	arr = make([]int, 27)

	fmt.Fscan(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &char)
		fmt.Fprintf(writer, "%d\n", char)

		if char != preChar {
			if arr[char-'a'] > 0 {
				arr[char-'a'] = -1
				resultCnt--
			} else if arr[char-'a'] != -1 {
				arr[char-'a']++
				resultCnt++
			}
		} else {
			if arr[char-'a'] != -1 && arr[char-'a'] == 0 {
				arr[char-'a']++
				resultCnt++
			}
		}

		preChar = char
	}

	fmt.Fprintf(writer, "%d", resultCnt)
}
