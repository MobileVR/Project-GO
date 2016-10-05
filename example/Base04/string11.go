package main

import "testing"

/*
除类型转换外，动态创建字符串也容易造成性能问题。
用加法操符字符拼接字符串时，每次都须重新分配内存。如此，在构建“超大”字符串时，性能就显得极差。
改进方法见 string12.go
 */
func main() {
}

/*
输出示例：
BenchmarkTest-4 10000 226285 ns/op  530348 B/op  999 allocs/op
 */
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test11()
	}
}

func test11() string {

	var s string

	for i := 0; i < 1000; i++ {
		s += "a"
	}

	return s
}
