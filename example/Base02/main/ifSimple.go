package main

func main() {

	/*
	if...else...
	条件表达式值必须是布尔类型，可省略括号，且左花括号不能另起一行。
	 */
	//if1()

	/*
	比较特别的是对初始化语句的支持，可定义块局部变量或执行初始化函数。
	局部变量的有效范围包含整个 if/else 块
	 */
	//if2()

}

/*
if...else...
条件表达式值必须是布尔类型，可省略括号，且左花括号不能另起一行。
 */
func if1() {

	x := 3

	if x > 5 {
		println("a")
	} else if x < 5&&x > 0 {
		println("b")
	} else {
		println("c")
	}

}

/*
比较特别的是对初始化语句的支持，可定义块局部变量或执行初始化函数。
局部变量的有效范围包含整个 if/else 块
 */
func if2() {

	x := 10

	/*if xinit(); x == 0 {//优先执行 xinit 函数
		println("a")
	}*/

	if a, b := x + 1, x + 10; a < b {
		//定义一个或多个局部变量（也可以是函数返回值）
		println("a")
	} else {
		println("b")
	}

}
