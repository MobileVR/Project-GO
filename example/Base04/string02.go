package main

/*
字符串默认值不是nil，而是""。
 */
func main() {

	var s string

	println(s)//""

	println(s == "")//true

	//println(s == nil)//错误：invalid operation: s == nil (mismatched types string and nil)

}
