package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
以切片语法（起始和结束索引号）返回字串时，其内部依旧指向原字节数组。
 */
func main() {

	s := "abcdefg"

	s1 := s[:3]//从头开始，仅指定结束索引位置
	s2 := s[1:4]//指定开始和结束位置，返回[start,end]
	s3 := s[2:]//指定开始位置，返回后面全部内容

	println(s1)//abc
	println(s2)//bcd
	println(s3)//cdefg

	/*
	提示：
	reflect.StringHeader 和 string 头结构相同
	unsafe.Pointer 用于指针类型转换
	我们可以从打印结果中看到子串和原字符串的内存地址是一样的：
	&reflect.StringHeader{Data:0x4b198a, Len:7}
	&reflect.StringHeader{Data:0x4b1aca, Len:3}
	都是Data:0x4b1aca。
	 */
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))//&reflect.StringHeader{Data:0x4b198a, Len:7}
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))//&reflect.StringHeader{Data:0x4b1aca, Len:3}

}
