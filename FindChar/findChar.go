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

	var S string
	var ret [26]int
	for i := 0; i < len(ret); i++ {
		ret[i] = -1
	}

	fmt.Fscanf(reader, "%s", &S)

	for i := 0; i < len(S); i++ {
		for j := 0; j < len(ret); j++ {
			if ret[j] == -1 {
				if byte(S[i]) == byte(j+97) {
					ret[j] = i
					break
				}
			} else {
				continue
			}
		}
	}

	for i := 0; i < len(ret); i++ {
		fmt.Fprintf(writer, "%d ", ret[i])
	}
}
