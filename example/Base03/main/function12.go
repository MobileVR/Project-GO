package main

import "errors"

/*
返回值
 */
func main() {

	/*
	有返回值的函数，必须有明确的return终止语句。
	 */
	//i := a12()
	//println(i)//1

	/*
	除非有 panic（异常处理机制的一种），或者无 break 的死循环，则无须 return 终止语句。
	 */
	/*i = b12()
	println(i)*/

	/*
	借鉴自动态语言的多返回值模式，函数得以返回更多状态，尤其是 error 模式。
	多返回值列表必须使用括号。
	 */
	/*i, err := c12(1, 0)
	if err != nil {
		println(i, "出错了")//0 出错了
	}*/

}

/*
有返回值的函数，必须有明确的return终止语句。
 */
func a12() int {
	return 1;
}

/*
除非有 panic（异常处理机制的一种），或者无 break 的死循环，则无须 return 终止语句。
 */
/*func b12() int {//missing return at end of function
	for {
		break
	}
}*/

/*
借鉴自动态语言的多返回值模式，函数得以返回更多状态，尤其是 error 模式。
多返回值列表必须使用括号。
 */
func c12(x, y int) (int, error) {
	return 0, errors.New("出错了,by zero");
}
