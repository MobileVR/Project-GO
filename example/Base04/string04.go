package main

/*
字符串支持 "!=、==、<、>、+、+="操作符。
 */
func main() {

	s := "ab" + //跨行时，加法操作符必须在上一行结尾
		"cd"

	println(s)//abcd

	println(s == "abcd")//true

	println(s > "abc")//true

}
