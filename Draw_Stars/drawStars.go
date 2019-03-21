package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var strArr = [3]string{"  *  ", " * * ", "*****"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	//N = 3 * 2^k
	var N int
	var whitePaper []string

	fmt.Fscan(reader, &N)
	whitePaper = make([]string, N)

	for K := 0; 3*int(math.Pow(2, float64(K))) <= N; K++ {
		drowStars(whitePaper, K, N)
	}

	for i := 0; i < len(whitePaper); i++ {
		fmt.Fprintf(writer, "%s\n", whitePaper[i])
	}

	writer.Flush()
}

func drowStars(whitePaper []string, K int, N int) {
	for i := 0; i < N; i++ {
		var space = ""

		for j := 0; j < 3*K; j++ {
			space += " "
		}

		for j := 0; j < len(strArr); j++ {
			whitePaper[i] = space + strArr[j]
		}
	}
}
