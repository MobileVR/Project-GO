package main

/*
panic，recover
与error相比，panic/recover 在使用方法上更接近 try/catch 结构化异常。
func panic(v interface{})
func recover() interface{}
比较有趣的是，它们是内置函数而非语句。panic 会立即中断当前函数流程，执行延迟调用。
而在延迟调用函数中，recover可捕获并返回panic提交的错误对象。
因为 panic 参数是空接口类型，因此可使用任何对象作为错误状态。
而 recover 返回结果同样要做转型才能获得具体信息。
建议：
除非是不可恢复性、导致系统无法正常工作的错误，否则不建议使用panic。
例如：文件系统没有操作权限，服务端口被占用，数据库未启动等情况。
 */
/*func main() {

	defer func() {
		if err := recover(); err != nil {
			//捕获错误
			log.Fatalln(err)
		}
	}()

	panic("i am dead")//引发错误
	println("exit.")//永不会执行

}*/

/*
无论是否执行 recover，所有延迟调用都会被执行。
但中断性错误会沿调用堆栈向外传递，要么被外层捕获，要么导致进程崩溃。
 */
/*func main() {

	defer func() {
		log.Println(recover())
	}()


	//输出示例：
	//test.2
	//test.1
	//2016/09/24 19:46:52 i am dead

	test22()

}

func test22() {

	defer println("test.1")
	defer println("test.2")

	panic("i am dead")

}*/

/*
连续调用 panic，仅最后一个会被 recover 捕获。
 */
/*func main() {

	//输出示例：
	//2016/09/24 19:53:05 you are dead
	//2016/09/24 19:53:05 fatal

	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal")
			}
		}
	}()

	defer func() {
		//类似重新抛出异常（rethrow）
		//可先recover捕获，包装后重新抛出
		panic("you are dead")
	}()

	panic("i am dead")

}*/

/*
在延迟函数中 panic，不会影响后续延迟调用执行。
而 recover 之后 panic，可被再次捕获。
另外，recover 必须在延迟调用函数中执行才能正常工作。
 */
/*
func main() {

	defer catch()//捕获
	defer log.Panicln(recover())//失败
	defer recover()//失败

	panic("i am dead")

}

func catch() {
	log.Panicln("catch:", recover())
}
*/

/*
考虑到 recover 特性，如果要保护代码片段，那么只能将其重构为函数调用。
 */
/*func main() {
	test22(5, 0)
}

func test22(x, y int) {

	z := 0

	func() {
		//利用匿名函数保护 "z = x / y"
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()

		z = x / y

	}()

	println("x / y =", z)//x / y = 0

}*/

/*
调试阶段，可使用runtime/debug.PrintStack函数输出完整调用堆栈信息。
 */
/*
func main() {

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	test22()

}

func test22() {
	panic("i am dead")
}
*/

