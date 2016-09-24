package main

/*
Go对参数的处理偏向保守，不支持默认值的可选参数，不支持命名实参。
调用时，必须按签名顺序传递类型和数量的实参，就算以"_"命名的参数也不能忽略。
在参数列表中，相邻的同类型参数可合并。
 */
func main() {

	function6(1, 2, "abc")//not enough arguments in call to function6（没有足够的参数调用function6方法，_参数不可省略）

}

func function6(x, y int, s string, _ bool) *int {
	return nil
}

/*
参数可视作函数局部变量，因此不能在相同层次定义同名变量。
形参是指函数定义中的参数，实参则是函数调用时所传递的参数。
形参类似函数局部变量，而实参则是函数外部对象，可以是常量、变量、表达式或函数等。
 */
func add(x, y int) int {
	x := 100//no new variables on left side of :=
	var y int//y redeclared in this block
	return x + y
}
