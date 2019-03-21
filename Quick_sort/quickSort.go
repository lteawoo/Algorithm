package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()

	var N, X int

	fmt.Scan(&N, &X)

	A := make([]int, N, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	quickSort(A, 0, N-1)

	for i := 0; i < len(A); i++ {
		fmt.Fprintf(buf, "%d ", A[i])
	}
}

func partition(arr []int, startNum, endNum int) int {
	left := startNum
	right := endNum
	pivot := (startNum + endNum) / 2

	//fmt.Printf("[pos]left : %d, right : %d, pivot : %d\n", left, right, pivot)
	//fmt.Printf("[dat]left : %d, right : %d, pivot : %d\n", arr[left], arr[right], arr[pivot])

	for left < right {
		for arr[left] <= arr[pivot] && left < right && left < endNum {
			left++
		}
		//fmt.Printf("[left calc] left : %d, right : %d, pivot : %d\n", left, right, pivot)
		//fmt.Printf("[left data]left : %d, right : %d, pivot : %d\n", arr[left], arr[right], arr[pivot])

		for arr[right] >= arr[pivot] && left < right && right > startNum {
			right--
		}
		//fmt.Printf("[right calc] left : %d, right : %d, pivot : %d\n", left, right, pivot)
		//fmt.Printf("[right data]left : %d, right : %d, pivot : %d\n", arr[left], arr[right], arr[pivot])

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
			/*
				for i := 0; i < len(arr); i++ {
					fmt.Printf("%d ", arr[i])
				}
				fmt.Printf("\n\n")
			*/
		}
	}

	arr[pivot], arr[right] = arr[right], arr[pivot]
	/*
		fmt.Printf("partition : ")
		for i := 0; i < len(arr); i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Printf("\n\n")
	*/
	return right
}

func quickSort(arr []int, startNum, endNum int) {
	if startNum < endNum {
		partition(arr, startNum, endNum)

		quickSort(arr, startNum, n-1)
		quickSort(arr, n+1, endNum)
	}
}
