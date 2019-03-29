문자열 반복
=============
<https://www.acmicpc.net/problem/2675>
# 문제
문자열 S를 입력받은 후에, 각 문자를 R번 반복해 새 문자열 P를 만든 후 출력하는 프로그램을 작성하시오. 즉, 첫 번째 문자를 R번 반복하고, 두 번째 문자를 R번 반복하는 식으로 P를 만들면 된다. S에는 QR Code "alphanumeric" 문자만 들어있다.

QR Code "alphanumeric" 문자는 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ\$%*+-./: 이다.
# 입력
첫째 줄에 테스트 케이스의 개수 T(1 ≤ T ≤ 1,000)가 주어진다. 각 테스트 케이스는 반복 횟수 R(1 ≤ R ≤ 8), 문자열 S가 공백으로 구분되어 주어진다. S의 길이는 적어도 1이며, 20글자를 넘지 않는다. 
# 출력
각 테스트 케이스에 대해 P를 출력한다.
# 풀이
<pre>
<code>
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
</code>
</pre>
첫 줄에 테스트케이스 개수, 둘째 줄부터 반복횟수 문자열 형식으로 입력받는다.

단순히 반복횟수와 문자열을 입력 받으면 문자열의 첫번째 문자부터 반복횟수만큼 출력하여 문제를 해결한다.