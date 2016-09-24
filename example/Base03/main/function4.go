package main

/*
函数只能判断其是否为 nil，不支持其他比较操作。
 */
func main() {
	println(a == nil)
	println(a == b)//无效操作：invalid operation: a == b (func can only be compared to nil)
}

func a() {
}

func b() {
}


