package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader
var writer *bufio.Writer

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	var n int
	var str string

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &str)
		fmt.Fprintln(writer, calc(str))
	}

	writer.Flush()
}

func calc(str string) int {
	var cnt, sum int

	for i := range str {
		if str[i] == 'O' {
			cnt++
			sum += cnt
		} else {
			cnt = 0
		}
	}

	return sum
}
