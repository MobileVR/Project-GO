package main

/*
匿名函数
匿名函数是指没有定义名字符号的函数。
除没有名字外，匿名函数和普通函数完全相同。
最大区别是，我们可在函数内部定义匿名函数，形式类似嵌套效果。
匿名函数可直接调用，保存到变量，作为参数或返回值。
1.将匿名函数赋值给变量，与为普通函数提供名字标识符有着根本的区别。
当然，编译器会为匿名函数生成一个“随机”符号名。
2.除闭包因素外，匿名函数也是一种常见重构手段。
可将大函数分解成多个相对独立的匿名函数块，然后用相对简洁的调用完成逻辑流程，以实现框架和细节分离。
相比语句块，匿名函数的作用域被隔离（不使用闭包），不会引发外部污染，更加灵活。
没有定义顺序限制，必要时可抽离，便于实现干净、清晰的代码层次。
 */
func main() {

	//Hello Word!
	/*func(s string) {
		println(s)
	}("Hello Word!")*/

	//赋值给变量
	/*add := func(x, y int) int {
		return x + y;
	}
	println(add(1, 2))//3*/

	//作为参数
	/*test15(func() {
		println("Hello Word!")
	})*/

	/*
	作为返回值
	 */
	//add := div15()
	//println(add(1, 2))//3

	/*
	普通函数和匿名函数都可作为结构体字段，或经通道传递
	 */
	//testStruct15()
	//testChannel15()

	/*
	不曾使用的匿名函数会被编译器当作错误
	 */
	func(s string) {
		//func literal evaluated but not used
		println(s)
	}//此处并未调用

}

func test15(f func()) {
	f()
}

func div15() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

/*
普通函数和匿名函数都可作为结构体字段，或经通道传递。
 */
func testStruct15() {

	type calc struct {
				       //定义结构体类型
		mul func(x, y int) int //函数类型字段
	}

	x := calc{
		mul:func(x, y int) int {
			return x + y
		},
	}

	println(x.mul(2, 3))
}

/*
普通函数和匿名函数都可作为结构体字段，或经通道传递。
 */
func testChannel15() {

	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}

	//println(<-c(1, 2))//错误：cannot call non-function c (type chan func(int, int) int)
	println((<-c)(1,
	2))

}