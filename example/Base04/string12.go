package main

import (
	"strings"
	"bytes"
)

/*
承接上文 string11.go
改进思路是预分配足够的内存空间。
常用方法是用 strings.Join 函数，它会统计所有参数长度，并一次性完成内存分配操作。
 */
func main() {
	println(test13())
}

/*
性能测试输出示例：
BenchmarkTest-4 100000 14868 ns/op	2048 B/op	2 allocs/op
编译器对 s1+s2+s3 这类表达式的处理方式和 strings.Join 类似。
 */
func test13() string {

	s := make([]string, 100)//分配足够的内存，避免中途扩张底层数组。

	for i := 0; i < 100; i++ {
		s[i] = "a"
	}

	return strings.Join(s, "")

}

/*
显然，改进后的算法有巨大的提升。另外， bytes.Buffer 也能完成类似错做，且性能相当。

性能测试示例：
BenchmarkTest-4 100000	15063 ns/op	2160 B/op	3 allocs/op

对于数量较少的字符串格式化拼接，可使用 fmt.Sprintf、text/template等方法。
注意：
字符串操作通常在堆上分配内存，这会对 Web 等高并发应用会造成较大影响，会有大量字符串对象要做垃圾回收。
建议使用 []byte 缓存池，或在栈上自行拼装等方式来实现 zero-garbage。
 */
func test12Buffer() string {

	var b bytes.Buffer

	//分配缓冲区大小，事先准备足够的内存，避免中途扩张
	b.Grow(1000)

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}

	return b.String()

}