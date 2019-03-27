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

	var N int

	fmt.Fscanf(reader, "%c", &N)
	fmt.Fprintf(writer, "%d", N)
}
