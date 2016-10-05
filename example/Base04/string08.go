package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

/*
要修改字符串，须将其转换为可变类型（[]rune或[]byte），待完成后再转换回来。
但不管如何转换，都须重新分配内存，并复制数据。
 */
func main() {

	s := "hello,world!"
	pp("s:%x\n", &s)//s:4b22ee

	bs := []byte(s)
	s2 := string(bs)
	pp("string to []byte,bs:%x\n", &bs)//string to []byte,bs:c04203a210
	pp("[]byte to string,s2:%x\n", &s2)//[]byte to string,s2:c04203a230

	rs := []rune(s)
	s3 := string(rs)
	pp("string to []rune,rs:%x\n", &rs)//string to []rune,rs:c04205c030
	pp("[]rune to string,s3:%x\n", &s3)//[]rune to string,s3:c042006310

	/*
	某些时候，转换操作会拖累算法性能，可尝试用“非安全”方法进行改善。
	 */
	t1 := []byte("hello,world!")
	t2 := toString(t1)

	pp("t1:%x\n", &t1)//t1:c04203a280
	pp("t2:%x\n", &t2)//t2:c04203a280

}

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

/*
该方法利用了[]byte和string头结构“部分相同”，以非安全的指针类型转换来实现类型“变更”，
从而避免了底层数组复制。在很多 Web Framework 中都能看到此类做法，在高并发压力下，
此种做法能有效改善执行性能。只是使用 unsafe 存在一定的风险，须小心谨慎！
 */
func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
