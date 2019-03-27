package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	//StudentCnt : 학생의 수
	StudentCnt = 5
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, avg int
	for i := 0; i < StudentCnt; i++ {
		fmt.Fscan(reader, &N)

		if N < 40 {
			N = 40
		}

		avg += N / StudentCnt
	}

	fmt.Fprintf(writer, "%d", avg)
}
