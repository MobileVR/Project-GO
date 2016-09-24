package main

import (
	"errors"
	"log"
)

/*
官方推荐的标准做法是返回 error 状态。
func Scanln(a ...interface{}) (n int, err error)

标准库将error定义为接口类型，以便实现自定义错误类型。
type error interface{
	Error() string
}

按惯例， error 总是最后一个返回参数。标准库提供了相关的创建函数，
可方便地创建包含简单错误文本的 error 对象。
 */
func main() {

	z, err := div20(5, 0)

	if err != nil {
		log.Fatalln(err)
	}

	println(z)
}

/*
	错误变量通常以 err 作为前缀，且字符串内容全部小写，没有结束标点，以便于嵌入到其他格式化字符串中输出。
	全局错误变量并非没有问题，因为它们可被用户重新赋值，这就可能导致结果不匹配。
不知道以后是否会出现只读变量功能，否则就只能靠自觉了。
	与 errors.New 类似的还有 fmt.Errorf，它返回一个格式化内容的错误对象。
 */
var errDivByZero = errors.New("division by zero")

func div20(x, y int) (int, error) {

	if y == 0 {
		return 0, errDivByZero
	}

	return x / y, nil

}
