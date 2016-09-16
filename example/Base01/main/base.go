package main

import "fmt"

/**
基本定义
 */
func main() {

	//基本定义
	//method1()

	//相同类型和不相同类型的多变量
	//method2()

	//以组的形式多行定义变量
	//method3()

	//使用浮点数时，须注意小数位的有效精度，相关细节可参考IEEE-754标准
	//method4()

	//别名类型无须转换，可直接赋值
	//method5()

	//拥有相同底层结构的并不表示就属于别名；就算在64位平台上 int 和 int64 结构完全一致，也分属不同类型，须显式转换。
	method6()
}

func method1() {

	//定义一个int32类型的变量，默认初始值为0
	var x int32

	//定义并初始化一个字符串
	var s = "hi"

	//打印
	println(x, s)

}

func method2() {

	//相同类型的多个变量
	var x, y int

	println(x, y)

	//不同类型的多个变量初始化值
	var a, b, c = 100, false, "say hi"

	println(a, b, c)

}

func method3() {

	//依照惯例，建议以组方式整理多行变量定义
	var (
		x, y int
		a, b = 100, "say hi"
	)

	println(x, y, a, b)

}

/**
默认浮点数类型是 float64
 */
func method4() {

	var a float32 = 1.1234567899        //注意：默认浮点数类型是 float64
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	println(a, b, c)//+1.123457e+000 +1.123457e+000 +1.123457e+000
	println(a == b, a == c, b == c)//true true true
	fmt.Printf("%v %v,%v\n", a, b, c)//1.1234568 1.1234568,1.1234568

}

/**
别名类型无须转换，可直接赋值
 */
func method5() {

	//byte alias for uint8
	//rune alias for int32

	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	ptf(c);

}
func ptf(x byte) {
	println(x)
}

/**
拥有相同底层结构的并不表示就属于别名；就算在64位平台上 int 和 int64 结构完全一致，也分属不同类型，须显式转换。
 */
func method6() {
	var x int = 100
	var y int64 = x	//cannot use x (type int) as type int64 in assignment
	add(x, y)	//cannot use y (type int64) as type int in argument to add
}
func add(x, y int) int {
	return x + y
}

