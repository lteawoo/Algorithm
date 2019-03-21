package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	max = 10000
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	var arr []bool

	arr = make([]bool, max)

	for i := 0; i < max; i++ {
		calc(arr, i+1, max)
	}

	for i := 0; i < max; i++ {
		if !arr[i] {
			fmt.Fprintf(writer, "%d\n", i+1)
		}
	}

	writer.Flush()
}

func calc(arr []bool, n int, max int) {
	var sum = n
	var temp = n

	for temp > 0 {
		//fmt.Printf("sum : %d + %d\n", sum, temp%10)
		sum += temp % 10
		temp /= 10
	}

	if sum <= max {
		arr[sum-1] = true
		//fmt.Printf("%d\n", sum)
		//calc(arr, sum, max)
	}

}
