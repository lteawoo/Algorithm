그룹 단어 체커
=============
<https://www.acmicpc.net/problem/1316>
# 문제
그룹 단어란 단어에 존재하는 모든 문자에 대해서, 각 문자가 연속해서 나타나는 경우만을 말한다. 예를 들면, ccazzzzbb는 c, a, z, b가 모두 연속해서 나타나고, kin도 k, i, n이 연속해서 나타나기 때문에 그룹 단어이지만, aabbbccb는 b가 떨어져서 나타나기 때문에 그룹 단어가 아니다.

단어 N개를 입력으로 받아 그룹 단어의 개수를 출력하는 프로그램을 작성하시오.
# 입력
첫째 줄에 단어의 개수 N이 들어온다. N은 100보다 작거나 같은 자연수이다. 둘째 줄부터 N개의 줄에 단어가 들어온다. 단어는 알파벳 소문자로만 되어있고 중복되지 않으며, 길이는 최대 100이다.
# 출력
첫째 줄에 그룹 단어의 개수를 출력한다.
# 풀이
<pre>
<code>
var n int
var str string

fmt.Fscan(reader, &n)

for i := 0; i < n; i++ {
    fmt.Fscan(reader, &str)
    fmt.Fprintln(writer, calc(str))
}
</code>
</pre>
main 함수의 내용, 첫번째 줄에는 테스트 케이스가 주어지므로 n변수에 입력받는다, 두번째 줄부터는 문자열을 케이스 별로 입력받고 바로 다음 calc 함수를 호출하여 로직을 해결한다.
* * *
<pre>
<code>
func calc(str string) int {
	var cnt, sum int

	for i := range str {
		if str[i] == 'O' {
			cnt++
			sum += cnt
		} else {
			cnt = 0
		}
	}

	return sum
}
</code>
</pre>
calc 함수에 알고리즘을 작성하였다.
문자열을 입력 받아 각 자리별로 O 문자이면 cnt를 증가시켜 sum에 더한다.
O가 아니면(X 이면) cnt를 0으로 치환하여 문제의 조건(연속된 0의 개수가 더해지는 점수)을 달성한다.