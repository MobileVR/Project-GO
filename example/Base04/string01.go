package main

import "fmt"

/*

字符串是不可变字节（byte）序列，其本身是一个复合结构。

type stringStruct struct{
	str unsafe.Pointer
	len int
}

头部指针指向字节数组，但没有NULL结尾。
默认以 UTF-8 编码存储 Unicode 字符，字面量里允许使用十六位进制、八进制和UTF编码格式。

 */
func main() {

	s := "hello word\x61\142\u0041"

	fmt.Printf("%s\n", s)//hello wordabA

	/*
	内置函数 len 返回字节数组长度， cap 不接受字符串类型参数。
	 */
	fmt.Printf("% x,len: %d\n", s, len(s))//68 65 6c 6c 6f 20 77 6f 72 64 61 62 41,len: 13

}
