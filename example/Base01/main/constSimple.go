package main

import (
	"unsafe"
	"fmt"
)

/**
常量
 */
func main() {

	//普通形式定义常量
	//const1()

	//以组的形式定义常量
	//const2()

	//可在函数代码块中定义常量，不曾使用的常量不会引发编译错误
	//const3()

	//如果显式指定类型，必须确保常量左右值类型一致，需要时可做显式转换。
	//右值不能超出常量类型取值范围，否则会引发溢出错误。
	//const4()

	//常量值也可以是某些编译器能计算出结果的表达式，如unsafe.Sizeof、len、cap等
	//const5()

	//在常量组中如不指定类型和初始值，则与上一行非空常量右值（表达式文本）相同
	//const6()

	//常量展开：常量不同于变量在运行期分配存储内存（非优化状态），常量通常会被编译器在预处理阶段直接展开，作为数据指令使用。
	//const7()

	//常量的两种状态对编译器的影响
	//const8()

}

func const1() {

	const x, y int = 123, 0x123

	const c = '我'

	println(x, y, c)//123 291 25105

}

func const2() {
	const (
		i, f = 1, 0.123
		b = false
	)

	println(i, f, b)//1 +1.230000e-001 false
}

/**
可在函数代码块中定义常量，不曾使用的常量不会引发编译错误
 */
func const3() {

	const x = 123

	println(x)//123

	const y = 1.23//未使用，不会引发编译错误

	{
		const x = "abc"//在不同作用域定义同名常量
		println(x)//abc
	}

}

/**
如果显式指定类型，必须确保常量左右值类型一致，需要时可做显式转换。
右值不能超出常量类型取值范围，否则会引发溢出错误。
 */
func const4() {

	const (
		x, y int = 99, -999
		b byte = byte(x)//x被指定为int类型，须显式转换为byte类型
		//n = uint8(y)//错误：constant -999 overflows uint8
	)

	//println(x, y, b, n)

}

/**
常量值也可以是某些编译器能计算出结果的表达式，如unsafe.Sizeof、len、cap等
 */
func const5() {
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello,world")
	)

	println(ptrSize, strSize)//8 11
}

/**
在常量组中如不指定类型和初始值，则与上一行非空常量右值（表达式文本）相同
 */
func const6() {

	const (
		x uint16 = 120
		y        //与上一行 x 类型、右值相同
		s = "abc"
		z        //与 s 类型、右值相同
	)

	//输出类型和值
	fmt.Printf("%T,%v\n", y, y)//uint16,120
	fmt.Printf("%T,%v", z, z)//string,abc

}

/**
常量不同于变量在运行期分配存储内存（非优化状态），常量通常会被编译器在预处理阶段直接展开，作为数据指令使用。
 */
func const7() {

	var x = 100
	const y = 100

	println(&x, x)
	println(&y, y)//cannot take the address of y

}

/**
常量的两种状态对编译器的影响
 */
func const8() {
	const x = 100		//无类型声明的常量
	const y byte = x        //直接展开 x，相当于 const y byte = 100

	const a int = 100        //显式指定常量类型，编译器会做强类型检查
	const b byte = a        //错误：cannot use a (type int) as type byte in const initializer

	println(x, y, a, b)
}
