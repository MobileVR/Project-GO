package main

import "fmt"

/*
不管是指针、引用类型，还是其他类型参数，都是值拷贝（pass-by-value）。
区别无非是拷贝目标对象，还是拷贝指针而已。
在函数调用前，会为形参和返回值分配内存空间，并将实参拷贝到形成内存。
 */
func main() {

	a := 0x100
	p := &a

	//输出实参p的地址
	fmt.Printf("pointer:%p,target:%v\n", &p, p)//pointer:0xc042024020,target:0xc042006280

	function7(p)

}

func function7(x *int) {
	//输出形参x的地址
	fmt.Printf("pointer:%p,target:%v\n", &x, x)//pointer:0xc042024030,target:0xc042006280
}
