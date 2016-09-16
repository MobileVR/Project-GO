package main

import "fmt"

func main() {

	//使用关键字 type 定义用户自定义类型，包括基于现有基础类型创建，或者是结构体、函数类型等。
	//custom1()

	//和var const 类似，多个 type 定义可合并成组，可在函数或代码块内定义局部类型。
	//custom2()

	/*
	即便指定了基础类型，也只表明它们有相同底层数据结构，两者不存在任何关系，
	属完全不同的两种类型。除操作符外，自定义类型不会继承基础类型的其他信息（包括方法）。
	不能视作别名，不能隐式转换，不能直接用于比较表达式。
 	*/
	//custom3()

	/*
	容易被忽视的是 struct tag，它也属于类型组成部分，而不仅仅是元数据描述
 	*/
	//custom4()

	/**
	同样，函数的参数顺序也属签名组成部分
 	*/
	//custom5()

	/*
	未命名类型转换规则：
	1、所属类型相同。
	2、基础类型相同，且其中一个是未命名类型。
	3、数据类型相同，将双向通道赋值给单向通道，且其中一个为未命名类型。
	4、将默认值 nil 赋值给切片、字典、通道、指针、函数或接口。
	5、对象实现了目标接口。
	 */
	//custom6()

}

/*
使用关键字 type 定义用户自定义类型，包括基于现有基础类型创建，或者是结构体、函数类型等。
 */
func custom1() {

	type flags byte

	const (
		read flags = 1 << iota        //1 * 2^0
		write        // 1 * 2^1
		exec        // 1 * 2^2
	)

	println(read, write, exec)//1 2 4


}

/*
和var const 类似，多个 type 定义可合并成组，可在函数或代码块内定义局部类型。
 */
func custom2() {

	type (        //组
		user struct {
			//结构体
			name string
			age  uint8
		}

		event func(string) bool//函数类型
	)

	u := user{"Tom", 20}
	fmt.Println(u)//{Tom 20}

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	r := f("abc")//abc
	println(r)//true

}

/*
即便指定了基础类型，也只表明它们有相同底层数据结构，两者不存在任何关系，
属完全不同的两种类型。除操作符外，自定义类型不会继承基础类型的其他信息（包括方法）。
不能视作别名，不能隐式转换，不能直接用于比较表达式。
 */
func custom3() {

	type data int
	var d data = 10

	var x int = d//cannot use d (type data) as type int in assignment
	println(x)

	println(d == x)//invalid operation: d == x (mismatched types data and int)

}

/*
容易被忽视的是 struct tag，它也属于类型组成部分，而不仅仅是元数据描述
 */
func custom4() {

	var a struct {
		//匿名结构类型
		x int `x`
		s string `s`
	}

	var b struct {
		x int
		s string
	}

	b = a//cannot use a (type struct { x int "x"; s string "s" }) as type struct { x int; s string } in assignment

	fmt.Println(b)

}

/**
同样，函数的参数顺序也属签名组成部分
 */
func custom5() {

	var a func(int, string)
	var b func(string, int)

	b = a        //cannot use a (type func(int, string)) as type func(string, int) in assignment

	b("s", 1)
}

/*
未命名类型转换规则：
1、所属类型相同。
2、基础类型相同，且其中一个是未命名类型。
3、数据类型相同，将双向通道赋值给单向通道，且其中一个为未命名类型。
4、将默认值 nil 赋值给切片、字典、通道、指针、函数或接口。
5、对象实现了目标接口。
 */
func custom6() {

	type data [2]int
	var d data = [2]int{1, 2}//基础类型相同，右值为未命名类型

	fmt.Println(d)//[1 2]

	a := make(chan int, 2)
	var b chan <- int = a        //双向通道转换为单向通道，其中 b 为未命名类型
	b <- 2
	fmt.Println(a, b)//0xc04202e070 0xc04202e070

}
