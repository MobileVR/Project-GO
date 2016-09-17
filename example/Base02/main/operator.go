package main

import "fmt"

/*
表达式通常是求值代码，可作为右值或参数使用。
而语句完成一个行为，比如 if、for 代码块。
表达式可作为语句用，但语句不能当作表达式。
 */
func main() {

	/*
	除位移操作外，操作数类型必须相同。如果其中一个是无显式类型声明的常量，那么该常量操作数会自动转型。
	*/
	//operator1()

	/*
	做位移运算时，右值必须是无符号整数，或可以转换的无显式类型常量
	*/
	//operator2()

	/*
	如果是非常量位移表达式，那么会优先将无显式类型的常量左操作数转型。
	*/
	//operator3()

	/*
	自增、自减不再是运算符。只能作为独立语句，不能用于表达式。
	*/
	//operator4()

}

/*
除位移操作外，操作数类型必须相同。如果其中一个是无显式类型声明的常量，那么该常量操作数会自动转型。
 */
func operator1() {

	const v = 20//无显式类型声明的常量

	var a byte = 10

	b := v + a//v 自动转换为 byte/uint8 类型

	fmt.Printf("%T,%v\n", b, b)//uint8,30

	const c float32 = 1.2

	d := c + v//v 自动转换为 float32 类型

	fmt.Printf("%T,%v\n", d, d)//float32,21.2

}

/*
做位移运算时，右值必须是无符号整数，或可以转换的无显式类型常量
 */
/*func operator2() {

	b := 23//b 是有符号 int 类型变量
	x := 1 << b//无效操作：invalid operation: 1 << b (shift count type int, must be unsigned integer)
	println(x)

}*/

/*
如果是非常量位移表达式，那么会优先将无显式类型的常量左操作数转型。
 */
func operator3() {

	a := 1.0 << 3//常量表达式（包括常量展开）
	fmt.Printf("%T,%v\n", a, a)//int,8

	var s uint = 3

	//因为 b 没有提供类型，那么编译器通过1.0推断，显然无法对浮点数做位移操作
	//b := 1.0 << s//无效操作：invalid operation: 1 << s (shift of type float64)
	//fmt.Printf("%T,%v\n", b, b)

	var c int32 = 1.0 << s//自动将 1.0 转换为 int32 类型
	fmt.Printf("%T,%v\n", c, c)//int32,8

}

/*
自增、自减不再是运算符。只能作为独立语句，不能用于表达式。
 */
func operator4() {

	a := 1
	//++a//语法错误：unexpected++（不能前置）syntax error: unexpected ++, expecting }

	//if (a++)>1{//语法错误：unexpected++，expecting（语句不能作为表达式使用）syntax error: unexpected ++, expecting }
	//}

	p := &a
	*p++//相当于 （*p）++
	println(a)//2

}
