OX 퀴즈
=============
<https://www.acmicpc.net/problem/8958>
# 문제
"OOXXOXXOOO"와 같은 OX퀴즈의 결과가 있다. O는 문제를 맞은 것이고, X는 문제를 틀린 것이다. 문제를 맞은 경우 그 문제의 점수는 그 문제까지 연속된 O의 개수가 된다. 예를 들어, 10번 문제의 점수는 3이 된다.

"OOXXOXXOOO"의 점수는 1+2+0+0+1+0+0+1+2+3 = 10점이다.

OX퀴즈의 결과가 주어졌을 때, 점수를 구하는 프로그램을 작성하시오.
# 입력
첫째 줄에 테스트 케이스의 개수가 주어진다. 각 테스트 케이스는 한 줄로 이루어져 있고, 길이가 0보다 크고 80보다 작은 문자열이 주어진다. 문자열은 O와 X만으로 이루어져 있다.
# 출력
각 테스트 케이스마다 점수를 출력한다.
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