package main

/*
内存地址是内存中每个字节单元的唯一编号，
而指针则是一个实体。
指针会分配内存空间，相当于一个专门用来保存地址的整型变量。
1、取址运算符 “&” 用于获取对象地址
2、指针运算符 “*” 用于间接引用目标对象
3、二级指针 “**T” 如包含包名则写成 “*package.T”
 */
func main() {

	/*
	并非所有对象都能进行取地址操作，但变量总是能正确返回（addressable）。
	指针运算符为左值时，我们可更新目标状态；而为右值时则是获取目标状态。
	*/
	//pointer1()

	/*
	new 函数也可以为引用类型分配内存，但这是不完整创建。以字典（map）为例，
	它仅分配了字典类型本身（实际就是个指针包装）所需内存，并没有分配键值存储内存，
	也没有初始化散列桶等内部属性，因此它无法正常工作。
	 */
	//pointer2()

	/*
	指针类型支持相等运算符，但不能做加减法运算和类型转换。
	如果两个指针指向同一地址，或都为 nil，那么他们相等。
	可通过 unsafe.Pointer 将指针转换为 uintptr 后进行加减法运算，但可能会造成非法访问。
	Pointer类型 C 语言中的 void* 万能指针，可用来转换指针类型，它能安全持有对象或对象成员，
	但 uintptr 不行。后者仅是一种特殊类型，并不引用目标对象，无法阻止垃圾回收器回收对象内存。
	*/
	//pointer3()

	/*
	指针没有专门指向成员的 "->" 运算符，统一使用 "." 选择表达式
	*/
	//pointer4()

	/*
	零长度（zero-size）对象的地址是否相等和具体的实现版本有关，不过肯定不等于 nil。
	即便长度为0，可该对象依然是“合法存在”的，拥有合法内存地址，这与 nil 语义完全不同。
	在 runtime/malloc.go 里有个 zerobase 全局变量，所有通过 mallocgc 分配的零长度对象都使用该地址。
	不过该例子中，并未调用 mallocgc 函数。
	*/
	//pointer5()

}

/*
并非所有对象都能进行取地址操作，但变量总是能正确返回（addressable）。
指针运算符为左值时，我们可更新目标状态；而为右值时则是获取目标状态。
 */
func pointer1() {

	x := 10

	var p *int = &x//获取地址，保存到指针变量
	*p += 20//用指针间接引用，并更新对象

	//输出指针所存储的地址，以及目标对象
	println(p, *p)//0xc04202df20 30

}

/*
new 函数也可以为引用类型分配内存，但这是不完整创建。以字典（map）为例，
它仅分配了字典类型本身（实际就是个指针包装）所需内存，并没有分配键值存储内存，
也没有初始化散列桶等内部属性，因此它无法正常工作。
 */
func pointer2() {
	m := map[string]int{"a":1}
	println(m["a"])//1
	//println(&m["a"])//cannot take the address of m["a"]
}

/*
指针类型支持相等运算符，但不能做加减法运算和类型转换。
如果两个指针指向同一地址，或都为 nil，那么他们相等。
可通过 unsafe.Pointer 将指针转换为 uintptr 后进行加减法运算，但可能会造成非法访问。
Pointer类型 C 语言中的 void* 万能指针，可用来转换指针类型，它能安全持有对象或对象成员，
但 uintptr 不行。后者仅是一种特殊类型，并不引用目标对象，无法阻止垃圾回收器回收对象内存。
 */
/*func pointer3() {
	x := 10
	p := &x
	p++//invalid operation: p++ (non-numeric type *int)
	var p2 *int = p + 1//invalid operation: p + 1 (mismatched types *int and int)

	p2 = &x

	println(p == p2)
}*/

/*
指针没有专门指向成员的 "->" 运算符，统一使用 "." 选择表达式
 */
func pointer4() {

	a := struct {
		x int
	}{}

	a.x = 100

	p := &a
	p.x += 100//相当于 p->x += 100

	println(p.x)//200
}

/*
零长度（zero-size）对象的地址是否相等和具体的实现版本有关，不过肯定不等于 nil。
即便长度为0，可该对象依然是“合法存在”的，拥有合法内存地址，这与 nil 语义完全不同。
在 runtime/malloc.go 里有个 zerobase 全局变量，所有通过 mallocgc 分配的零长度对象都使用该地址。
不过该例子中，并未调用 mallocgc 函数。
*/
func pointer5() {
	var a, b struct{}

	println(&a, &b)//0xc04202df27 0xc04202df27
	println(&a == &b, &a == nil)//true false
}
