package main

func reverseString(s string) (string, bool) {
	str := []rune(s)
	length := len(str)
	if length > 5000 {
		return s, false
	}

	for i := 0; i < length/2; i++ {
		// str[i] 代表第 i 个字符
		// str[length-i-1] 代表第 length-i-1 个字符
		str[i], str[length-i-1] = str[length-i-1], str[i]
	}
	return string(str), true
}

func main() {
	str, flag := reverseString("hello")
	println(str, flag)
}
