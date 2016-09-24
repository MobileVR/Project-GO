package main

/*
从函数返回局部变量指针是安全的，编译器会通过逃逸分析（escape analysis）来决定是否在堆上分配内存。
函数内联（inline）对内存分配有一定的影响。如果该例中允许内联，那么就会直接在栈上分配内存。
当前编译器（GO1.6）并未实现尾递归优化（tail-call optimization）。尽管GO执行栈的上限是GB规模，
轻易不会出现堆栈溢出（stack overflow）错误，但依然需要注意拷贝栈的复制成本。
 */
func main() {

	var a *int = function5()

	println(a, *a)//0xc04202ff20 256

}

func function5() *int {
	a := 0x100
	return &a
}
