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

	var T, R int
	var S string

	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		fmt.Fscan(reader, &R, &S)

		for j := 0; j < len(S); j++ {
			for k := 0; k < R; k++ {
				fmt.Fprintf(writer, "%c", S[j])
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
