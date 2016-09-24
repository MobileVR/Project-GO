package main

import "errors"

/*
命名返回值
 */
func main() {

	/*
	命名返回值和参数一样，可当作函数局部变量使用，最后由 return 隐式返回。
	 */
	//z, err := div14(1, 0)
	//println(z, err)//0 (0x4961c0,0xc04202df18)

}

/*
对返回值命名和简短变量定义一样，优缺点共存。
 */
/*func paging14(sql string, index int) (count int, pages int, err error) {
}*/

/*
命名返回值和参数一样，可当作函数局部变量使用，最后由 return 隐式返回。
 */
func div14(x, y int) (z int, err error) {

	if y == 0 {
		err = errors.New("出错了，y为0了")
		return
	}

	z = x / y

	return//相当于 " return z,err "

}

/*
这些特殊的“局部变量”会被不同层级的同名变量遮蔽。好在编译器能检查到此类状况，只要改为显式 return 返回即可。
 */
func add14(x, y int) (z int) {
	{
		z := x + y//新定义的同名局部变量，同名屏蔽
		//return//错误：z is shadowed during return；解决办法：改成 " return z " 即可
		return z
	}

	return
}

/*
除遮蔽外，我们还必须对全部返回值命名，否则编译器会搞不清状况。
显然编译器在处理 return 语句的时候，会跳过未命名返回值，无法准确匹配。
 */
func test14() (int, s string, e error) {
	return 0, "", nil//cannot use 0 (type int) as type string in return argument
}

/*
如果返回值类型能明确其含义，就尽量不要对其命名。
 */
/*
func NewUser() (*User, error) {
}*/
