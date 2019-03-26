RGB 거리
=============
<https://www.acmicpc.net/problem/1149>
# 문제
RGB거리에 사는 사람들은 집을 빨강, 초록, 파랑중에 하나로 칠하려고 한다. 또한, 그들은 모든 이웃은 같은 색으로 칠할 수 없다는 규칙도 정했다. 집 i의 이웃은 집 i-1과 집 i+1이다.

각 집을 빨강으로 칠할 때 드는 비용, 초록으로 칠할 때 드는 비용, 파랑으로 드는 비용이 주어질 때, 모든 집을 칠할 때 드는 비용의 최솟값을 구하는 프로그램을 작성하시오.
# 입력
첫째 줄에 집의 수 N이 주어진다. N은 1,000보다 작거나 같다. 둘째 줄부터 N개의 줄에 각 집을 빨강으로 칠할 때, 초록으로 칠할 때, 파랑으로 칠할 때 드는 비용이 주어진다. 비용은 1,000보다 작거나 같은 자연수이다.
# 출력
첫째 줄에 모든 집을 칠할 때 드는 비용의 최솟값을 출력한다.
# 풀이
<pre>
<code>
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
</code>
</pre>
main 함수의 내용, 첫 집은 RGB 전부 연산 하여 최소 비용을 구해야 한다. 입력 받은 집의 개수가(N) 1이면 그저 첫 집의 최소만 구하면 되므로 해당 부분을 조건처리한다.
* * *
<pre>
<code>
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
</code>
</pre>
핵심은 N번째 집의 색상이 정해졌을 때 N+1번째 집의 이전 색상(prevPos)를 제외한 최소 비용 색은 고정이라는 점이다, 마지막 집의 최소 비용을 구하고, 그 이전 집은 그 최소 비용에 자신의 비용을 더하여 최소 비용을 구한다. 그리고 계속 거슬러 올라가 같은 연산을 반복하다 보면 첫번째 집이 선택한 색상의 최소비용이 구해지게 되는 것이다, 그리고 경우의 수가 많아 질 수록 같은 연산을 반복하게 되므로(N번째 집의 R 색상의 경우와 B 색상의 경우는 N+1번째 집의 G 색상의 최소비용을 구해야한다.) 한번 연산한 후 cache 배열에 저장하여 불필요한 재귀를 막는다.