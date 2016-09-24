package main

/*
延迟调用
语句 defer 向前函数注册稍后执行的函数调用。
这些调用被称作延迟调用，因为它们直到当前函数执行结束前才被执行，
常用于资源释放、解除锁定，以及错误处理等操作。
 */
/*func main() {

	f, err := os.Open("./main.go")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()//仅注册，直到main函数退出前才执行

}*/

/*
注意，延迟调用注册的是调用，必须提供执行所需参数（哪怕为空）。
参数值在注册的被复制并缓存起来。如对状态敏感，可改用指针或闭包。
延迟调用可修改当前函数命名返回值，但其自身返回值被抛弃。
 */
/*func main() {

	x, y := 1, 2

	defer func(a int) {
		//y为闭包引用，注册时复制调用参数
		println("defer x,y = ", a, y)//defer x,y =  1 202
	}(x)

	x += 100//对x的修改不会影响延迟调用
	y += 200

	println(x, y)//101 202

}*/

/*
多个延迟注册按FILO次序执行。
 */
/*func main() {
	//依次输出：
	// b
	// a
	defer println("a")
	defer println("b")
}*/

/*
编译器通过插入额外指令来实现延迟调用执行，而 return 和 panic 语句都会终止当前函数流程，引发延迟调用。
另外，return 语句不是 ret 汇编指令，它会优先更新返回值。
 */
/*func main() {
	println("test:", test17())//test: 200
}

func test17() (z int) {

	defer func() {
		println("defer:", z)//defer: 100
		z += 100//修改命名返回值
	}()

	return 100//实际执行次序：z=100;call defer;ret

}*/
