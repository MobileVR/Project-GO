package main

import "fmt"

/*
变参
 */
func main() {

	/*
	变参本质上就是一个切片。只能接收一到多个同类型参数，且必须放在列表尾部。
	 */
	//function11("abc", 1, 2, 3, 4)//[]int,[1 2 3 4]

	/*
	将切片作为变参时，须进行展开操作。如果是数组，先将其转换为切片。
	 */
	//a := [3]int{10, 20, 30}
	//test11(a)//cannot use a (type [3]int) as type int in argument to test11
	//test11(a[:])//cannot use a[:] (type []int) as type int in argument to test11
	//切换为slice后展开
	//test11(a[:]...)//[10 20 30]

	/*
	既然变参是切片，那么参数复制的仅是切片自身，并不包括底层数组，也因此可修改原数据。
	如果需要，可用内置函数 copy 复制底层数据。
	 */
	a := []int{10, 20, 30}
	fmt.Println(a)//[10 20 30]
	demo11(a...)
	fmt.Println(a)//[110 120 130]

}

/*
变参本质上就是一个切片。只能接收一到多个同类型参数，且必须放在列表尾部。
 */
func function11(s string, a ...int) {
	fmt.Printf("%T,%v\n", a, a)//显示类型和值
}

/*
将切片作为变参时，须进行展开操作。如果是数组，先将其转换为切片。
 */
func test11(a ...int) {
	fmt.Println(a)
}

/*
既然变参是切片，那么参数复制的仅是切片自身，并不包括底层数组，也因此可修改原数据。
如果需要，可用内置函数 copy 复制底层数据。
 */
func demo11(a ...int) {
	for i := range a {
		a[i] += 100
	}
}
