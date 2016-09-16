package main

/**
类型转换
 */
func main() {

	/**
	除常量、别名类型以及未命名类型外，Go 强制要求使用显式类型转换。
	加上不支持操作符重载，所以我们总是能确定语句及表达式的明确含义。
	 */
	//type1()

	//如果转换的目标是指针、单向通道或没有返回值的函数类型，那么必须使用括号，以避免造成语法分解错误。
	type2()

}

/**
除常量、别名类型以及未命名类型外，Go 强制要求使用显式类型转换。
加上不支持操作符重载，所以我们总是能确定语句及表达式的明确含义。
*/
func type1() {

	//混合类型表达式必须确保类型一致
	a := 10
	b := byte(a)
	c := a + int(b)
	println(a, b, c)//10 10 20

	//同样不能将非 bool 类型结果当作 true/false 使用
	//x := 100
	//var b bool = x//cannot use x (type int) as type bool in assignment
	//if x {
	//non-bool x (type int) used as if condition
	//}

}

/**
如果转换的目标是指针、单向通道或没有返回值的函数类型，那么必须使用括号，以避免造成语法分解错误。
 */
func type2() {

	x := 100
	//p := *int(&x)//cannot convert &x (type *int) to type int;invalid indirect of int(&x) (type int)
	p := (*int)(&x)//0xc04202df20
	println(p)

}
