음계
=============
<https://www.acmicpc.net/problem/2920>
# 문제
다장조는 c d e f g a b C, 총 8개 음으로 이루어져있다. 이 문제에서 8개 음은 다음과 같이 숫자로 바꾸어 표현한다. c는 1로, d는 2로, ..., C를 8로 바꾼다.

1부터 8까지 차례대로 연주한다면 ascending, 8부터 1까지 차례대로 연주한다면 descending, 둘 다 아니라면 mixed 이다.

연주한 순서가 주어졌을 때, 이것이 ascending인지, descending인지, 아니면 mixed인지 판별하는 프로그램을 작성하시오.
# 입력
첫째 줄에 8개 숫자가 주어진다. 이 숫자는 문제 설명에서 설명한 음이며, 1부터 8까지 숫자가 한 번씩 등장한다.
# 출력
첫째 줄에 ascending, descending, mixed 중 하나를 출력한다.
# 풀이
<pre>
<code>
var N []int
N = make([]int, COUNT)

fmt.Fprintf(writer, "%s", getScale(N))
</code>
</pre>
main 함수의 내용, 8개의 숫자를 받을 배열을 생성한다, 그 후 getScale을 호출하여 판별한 문자열을 받고 출력한다.
* * *
<pre>
<code>
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

</code>
</pre>
8개의 숫자를 받음과 동시에 연산한다, flag 변수를 이용하여 각 자리의 숫자를 받을 때 마다 1 증가하면 ascending, -1로 감소하면 descending, 그리고 외의 숫자로 증가되면 mixed로 판별하여 리턴한다.