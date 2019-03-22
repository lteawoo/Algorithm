package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	//COUNT is input max size
	COUNT = 8
)

var reader *bufio.Reader
var writer *bufio.Writer

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	var N []int
	N = make([]int, COUNT)

	fmt.Fprintf(writer, "%s", getScale(N))

	writer.Flush()
}

func getScale(N []int) string {
	var scaleStr string
	var flag int

	for i := 0; i < COUNT; i++ {
		fmt.Fscan(reader, &N[i])

		if i > 0 {
			if flag == 0 {
				flag = N[i-1] - N[i]

				if flag == -1 {
					scaleStr = "ascending"
				} else if flag == 1 {
					scaleStr = "descending"
				}
			} else if flag != N[i-1]-N[i] {
				scaleStr = "mixed"
			}
		}
	}

	return scaleStr
}
