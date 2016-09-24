package main

import (
	"sync"
	"testing"
)

var m sync.Mutex

/*
defer 性能
相比直接用 call 汇编指令调用函数，延迟调用须花费更大代价。
这其中包括注册、调用等操作，还有额外的缓存开销。
以最常用的 mutex 为例，我们简单对比一下两者的性能差异。
 */
func main() {

	//TODO 调用并执行
	/*
	输出示例：
	BenchmarkCall-4  100000000 22.4ns/op
	BenchmarkDefer-4 20000000  93.8ns/op
	相差几倍的结果足以引起重视。尤其是那些性能要求高且压力大的算法，应避免使用延迟调用。
	 */
}

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}