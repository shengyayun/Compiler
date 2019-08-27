package main

//是否是字母
func isAlpha(ch rune) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

//是否是数字
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

//是否是空白字符
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
