package main

/**
简短模式
 */
func main() {

	//简短模式
	//short1()

	//简短模式使用误区
	//short2()

	//退化赋值操作的前提条件是：最少有一个新变量被定义，且必须是同一作用域
	//short3()

	//退化赋值操作的前提条件是：最少有一个新变量被定义，且必须是同一作用域
	//short4()

	//不同作用域，全部是新变量定义
	//short5()

	//在处理函数错误返回值时，退化赋值允许我们重复使用err变量，这是相当有益的
	//short6()

}

/**
简单模式：short variable declaration
1.定义变量，同时显式初始化
2.不能提供数据类型
3.只能在函数内部使用
 */
func short1() {

	//在函数内部，可以省略var关键字
	x := 100

	a, b := false, "say hi"

	//打印
	println(x, a, b)

}

//全局变量 x
var x = 100

func short2() {

	println(&x, x)//0x496178 100

	x := "abc"//重新定义和初始化同名局部变量

	println(&x, x)//0xc04202df18 abc

}

func short3() {

	x := 100

	println(&x)//0xc04202df20

	//注意：x退化为赋值操作，仅有y是变量定义
	x, y := 200, "abc"

	println(&x, x)//0xc04202df20 200
	println(y)//abc

}

/**
退化赋值操作的前提条件是：最少有一个新变量被定义，且必须是同一作用域
 */
func short4() {

	x := 100
	println(&x)

	//x := 200//错误：no new variables on left side of :=
	println(&x)

}

/**
退化赋值操作的前提条件是：最少有一个新变量被定义，且必须是同一作用域
 */
func short5() {

	x := 100
	println(&x, x)//0xc04202df20 100

	{
		//不同作用域，全部是新变量定义
		x, y := 200, 300
		println(&x, x, &y, y)//0xc04202df18 200 0xc04202df10 300
	}

}

/**
在处理函数错误返回值时，退化赋值允许我们重复使用err变量，这是相当有益的
 */
func short6() {

	//f, err := os.Open("/dev/random")

	//...

	//buf := make([]byte, 1024)
	//n, err := f.Read(buf)//err退化赋值操作，n新定义

	//...

}
