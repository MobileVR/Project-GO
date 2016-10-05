package main

/*
考虑到字符串只读特征，转换时复制数据到新分配内存是可以理解的。
当然，性能同样重要，编译器会为某些场合进行专门优化，避免额外分配和复制操作：
1、将 []byte 转换为 string key，去 map[string] 查询的时候；
2、将 string 转换为 []byte ，进行 for range 迭代时，直接取字节赋值给局部变量。
 */
func main() {

	m := map[string]int{
		"abc":123,
	}

	key := []byte("abc")

	x, ok := m[string(key)]

	println(x, ok)//123 true

}
