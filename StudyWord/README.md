단어 공부
=============
<https://www.acmicpc.net/problem/1157>
# 문제
알파벳 대소문자로 된 단어가 주어지면, 이 단어에서 가장 많이 사용된 알파벳이 무엇인지 알아내는 프로그램을 작성하시오. 단, 대문자와 소문자를 구분하지 않는다.
# 입력
첫째 줄에 알파벳 대소문자로 이루어진 단어가 주어진다. 주어지는 단어의 길이는 1,000,000을 넘지 않는다.
# 출력
첫째 줄에 이 단어에서 가장 많이 사용된 알파벳을 대문자로 출력한다. 단, 가장 많이 사용된 알파벳이 여러 개 존재하는 경우에는 ?를 출력한다.
# 풀이
<pre>
<code>
var countArr []int
countArr = make([]int, 'Z'-'A'+1)

var max = 0
var maxChar byte
for {
	char, _ := reader.ReadByte()

	//문자열의 끝(10, 13 각각 Line Feed, Return Carriage)을 체크하기 위한.. 그외의 문자가 아닌것도 겸사겸사 체크하여 중지하던지 함
	if (('A' > char) || ('Z' < char)) && (('a' > char) || ('z' < char)) {
		break
	}
	//소문자면 대문자로 변환(문자가 'a' 이상 -> 97 이상 -> 소문자면)
	if char >= 'a' {
		char = char - ('a' - 'A')
	}

	idx := int(char - 'A')

	countArr[idx]++

	//최대 카운트 계산
	if max < countArr[idx] {
		max = countArr[idx]
		maxChar = char
	} else if max == countArr[idx] {
		maxChar = '?'
	}
}

fmt.Fprintf(writer, "%c", maxChar)
</code>
</pre>
카운트를 저장하는 알파벳 개수 만큼 배열을 생성, 문자열을 바이트 단위로 끊어서 연산한다

Go만 그러는지는 모르겠지만(C에서는 문자열의 끝에 null 이였던 것 같지만) 문자열의 끝엔 LF, RC가 온다 체크 후 break

소문자로 들어오면 32를 빼준다('a' - 'A') 그러면 대문자 아스키코드가 된다.

배열의 위치를 나타내는 변수 idx를 선언 후 문자에 'A'를 빼서 'A'가 0을 가리키게 끔 처리

동시에 max와 countArr[idx]와 비교하여 최대 값을 찾아낸다, max와 countArr[idx]가 같으면 '?'를 출력하기 위해 maxChar에 대입하여 알고리즘 완성~