package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	var N int
	var arr [][]int
	var cacheArr [][]int
	var ret []int
	var min int

	fmt.Fscan(reader, &N)

	arr = make([][]int, N)
	cacheArr = make([][]int, N)
	ret = make([]int, 3)

	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 3)
		cacheArr[i] = make([]int, 3)

		//각 비용 입력 받음
		for j := 0; j < len(arr[i]); j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		ret[i] = arr[0][i]
		if N > 1 {
			ret[i] += calc(arr, cacheArr, i, 1)
		}

		//최종 결과 최소값 구함
		if min == 0 || min > ret[i] {
			min = ret[i]
		}
	}

	fmt.Fprintf(writer, "%d", min)

	writer.Flush()
}

func calc(arr [][]int, cacheArr [][]int, prevPos int, row int) int {
	var min int

	for i := 0; i < 3; i++ {
		if prevPos != i {
			//cache 배열 값이 0이면 cache되지 않았으니 연산 값이 배열에 저장될 것이고
			//cache 배열을 조회하여 값이 있으면 재귀를 하지 않고 바로 진행하여 불필요한 중복 연산을 제거한다.
			if cacheArr[row][i] == 0 {
				cacheArr[row][i] = arr[row][i]

				//배열의 마지막인 경우는 재귀필요 X
				if row < len(arr)-1 {
					cacheArr[row][i] += calc(arr, cacheArr, i, row+1)
				}
			}

			//최소값 계산
			if min == 0 || min > cacheArr[row][i] {
				min = cacheArr[row][i]
			}
		}
	}

	return min
}
