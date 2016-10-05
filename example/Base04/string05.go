package main

/*
允许以索引号访问字节数组（非字符），但不能获取元素地址。
 */
func main() {

	s := "abc"

	println(s[1])//98

	println(&s[1])//错误：cannot take the address of s[1]

}
