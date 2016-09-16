package main

/**
引用类型
 */
func main() {

	/**
	内置函数 new 按指定类型长度分配零值内存，返回指针，并不关心类型内部构造和初始化方式。
	而引用类型则必须使用 make 函数创建，编译器会将 make 转换为目标类型专用的创建函数（或指令），以确保完成全部内存分配和相关属性初始化。
 	*/
	//reference1()

	/**
	new 函数也可以为引用类型分配内存，但这是不完整创建。以字典（map）为例，
	它仅分配了字典类型本身（实际就是个指针包装）所需内存，并没有分配键值存储内存，
	也没有初始化散列桶等内部属性，因此它无法正常工作。
	*/
	//reference2()

}

/**
内置函数 new 按指定类型长度分配零值内存，返回指针，并不关心类型内部构造和初始化方式。
而引用类型则必须使用 make 函数创建，编译器会将 make 转换为目标类型专用的创建函数（或指令），以确保完成全部内存分配和相关属性初始化。
 */
func reference1() {
	m := mkmap()
	println(m["a"])        //1

	s := mkslice()
	println(s[0])        //100
}

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s;
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

/**
new 函数也可以为引用类型分配内存，但这是不完整创建。以字典（map）为例，
它仅分配了字典类型本身（实际就是个指针包装）所需内存，并没有分配键值存储内存，
也没有初始化散列桶等内部属性，因此它无法正常工作。
 */
func reference2() {

	p := new(map[string]int)//函数 new 返回指针
	m := *p
	m["a"] = 1//panic: assignment to entry in nil map（运行期错误）
	println(m)

}