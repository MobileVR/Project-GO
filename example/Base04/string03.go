package main

/*
使用 "`" 定义不做转义处理的原始字符串（raw string），支持跨行。
编译器不会解析原始字符串内的注释语句，且前置缩进空格也属于字符串内容。
（有缩进的情况）例如：
	s:=`line\r\n,
	line2`
那么它打印出来的也是缩进的情况：
	s:=`line\r\n,
		line2`
同理，若不缩进：
	s:=`line\r\n,
line2`
那么它打印出来的就是不缩进的情况：
	s:=`line\r\n,
	line2`
两者的区别演示到此。
 */
func main() {

	s:=`line\r\n,
line2`

	/*
	输出示例：
	line\r\n,
	line2
	 */
	println(s)

}
