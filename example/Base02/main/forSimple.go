package main

import "fmt"

func main() {

	/*
	for 循环基本定义
	 */
	//for1()

	/*
	初始化语句仅执行一次。条件表达式中如有函数调用，须确认是否会重复执行。
	可能会被编译器优化掉，也可能是动态结果须每次执行确认。
	 */
	//for2()

	/*
	可用 for...range 完成数据迭代，支持字符串、数组、数组指针、切片、字典、通道类型
	返回索引、键值数据。
	注意：没有相关接口实现自定义类型迭代，除非基础类型是上述类型之一。
	 */
	//for3()

	/*
	允许返回单值，或用"_"忽略
	 */
	//for4()

	/*
	无论普通 for 循环，还是 range 迭代，其定义的局部变量都会重复使用。
	注意：这对闭包存在一些影响。
	 */
	//for5()

	/*
	注意：range会复制目标数据。受直接影响的是数组，可改用数组指针或切片类型。
	相关数据类型中，字符串、切片基本结构是个很小的结构体，而字典、通道本身是指针封装，复制成本很小，无须专门优化。
	 */
	//for6()

	/*
	如果range目标表达式是函数调用，也仅被执行一次。
	 */
	//for7()

	/*
	使用 goto 前，必须定义标签。标签区分大小写，且未使用的标签会引发编译错误。
	 */
	//for8()

	/*
	不能跳转到其他函数，或内层代码块内。
	 */
	//for9()

	/*
	break：用于switch、for、select语句，终止整个语句块执行。
	continue：仅用于for循环，终止后续逻辑，立即进入下一轮循环。
	 */
	//for10()

	/**
	配合标签，break 和 continue 可在多层嵌套中指定目标层级
	 */
	//for11()

}

/*
for 循环基本定义
 */
func for1() {

	//初始化表达式支持函数调用或定义局部变量
	/*for i := 0; i < 3; i++ {
		println(i)
	}*/

	//类似 "while x < 3{}" 或 "for ; x < 3; {}"
	/*x := 0
	for x < 3 {
		println(x)
		x++
	}*/

	//类似 "while true {}" 或 "for true {}"
	/*for {
		println("无限循环")
		break
	}*/

}

/*
初始化语句仅执行一次。条件表达式中如有函数调用，须确认是否会重复执行。
可能会被编译器优化掉，也可能是动态结果须每次执行确认。
 */
func for2() {

	//初始化语句 count 函数仅执行一次
	for i, c := 0, count(); i < c; i++ {
		println(i)
	}

	c := 0
	//条件表达式中的 count 重复执行
	//规避方式就是在初始化表达式中定义局部变量保存 count 结果。
	for c < count() {
		println(c)
		c++
	}

}

func count() int {
	println("count 执行了")
	return 3
}

/*
可用 for...range 完成数据迭代，支持字符串、数组、数组指针、切片、字典、通道类型
返回索引、键值数据。
注意：没有相关接口实现自定义类型迭代，除非基础类型是上述类型之一。
 */
func for3() {

	data := [3]string{"a", "b", "c"}

	/*
	0 a
	1 b
	2 c
	 */
	for i, s := range data {
		println(i, s)
	}

}

/*
允许返回单值，或用"_"忽略
 */
func for4() {

	data := [3]string{"a", "b", "c"}

	//0、1、2
	for i := range data {
		//只返回 lst value
		println(i)
	}

	//a、b、c
	for _, s := range data {
		//忽略 lst value
		println(s)
	}

	for range data {
		//仅迭代，不返回。可用来执行清空channel等操作
	}

}

/*
无论普通 for 循环，还是 range 迭代，其定义的局部变量都会重复使用。
注意：这对闭包存在一些影响。
 */
func for5() {

	data := [3]string{"a", "b", "c"}

	/*
	0xc04202fe98 0xc04202feb0
	0xc04202fe98 0xc04202feb0
	0xc04202fe98 0xc04202feb0
	 */
	for i, s := range data {
		println(&i, &s)
	}

}

/*
注意：range会复制目标数据。受直接影响的是数组，可改用数组指针或切片类型。
相关数据类型中，字符串、切片基本结构是个很小的结构体，而字典、通道本身是指针封装，复制成本很小，无须专门优化。
 */
func for6() {

	data := [3]int{10, 20, 30}

	for i, x := range data {
		//从data复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		/*
		x:10,data:110
		x:20,data:220
		x:30,data:330
		range 返回的依旧是复制值
		 */
		fmt.Printf("x:%d,data:%d\n", x, data[i])

	}

	for i, x := range data[:] {
		//仅复制slice，不包括底层array
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		/*
		x:110,data:210
		x:420,data:420
		x:630,data:630
		当 i==0 修改 data 时，x 已经取值，所以是110；复制的仅是 slice 自身，底层 array 依旧是原对象
		 */
		fmt.Printf("x:%d,data:%d\n", x, data[i])
	}

}

/*
如果range目标表达式是函数调用，也仅被执行一次。
 */
func for7() {

	/*
	data 执行了
	0 10
	1 20
	2 30
	 */
	for i, s := range data() {
		println(i, s)
	}

}
func data() []int {
	println("data 执行了")
	return []int{10, 20, 30}
}

/*
使用 goto 前，必须定义标签。标签区分大小写，且未使用的标签会引发编译错误。
 */
/*func for8() {

	start://label start defined and not used
	for i := 0; i < 3; i++ {
		println(i)
		if i > 1 {
			goto exit
		}
	}

	exit:
	println("exit.")

}*/

/*
不能跳转到其他函数，或内层代码块内。
 */
/*func for9() {
	for i := 0; i < 3; i++ {
		loop:
		println(i)
	}

	goto test//错误：label test not defined
	goto loop//错误：goto loop jumps into block
}

func test() {
	test:
	println("test")
	println("test exit")
}*/

/*
break：用于switch、for、select语句，终止整个语句块执行。
continue：仅用于for循环，终止后续逻辑，立即进入下一轮循环。
 */
func for10() {

	//1、3、5
	for i := 0; i < 10; i++ {

		if i % 2 == 0 {
			continue//跳过本次循环，进入下一轮循环
		}

		if i > 5 {
			break//结束整个循环
		}

		println(i)

	}
}

/**
配合标签，break 和 continue 可在多层嵌套中指定目标层级
 */
func for11() {

	/*
	0:0 0:1 0:2
	1:0 1:1 1:2
	2:0 2:1 2:2
	*/
	outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j > 2 {
				println()
				continue outer
			}

			if i > 2 {
				break outer
			}

			print(i, ":", j, " ")
		}
	}
}

