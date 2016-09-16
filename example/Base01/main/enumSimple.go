package main

/**
枚举类型
 */
func main() {

	//以iota来表达枚举类型
	//enum1()

	//enum枚举示例：取巧的办法
	//enum2()

	//自增作用范围为常量组。可在多常量定义中使用多个iota，它们各自单独计数，只须确保组中每行常量的列数量相同即可
	//enum3()

	//如中断 iota 自增，则必须显式恢复。且后续自增值按行序递增，而非 C enum 那般按上一取值递增
	//enum4()

	//自增默认类型为 int，可显式指定类型
	//enum5()

	/**
	在实际编码中，建议用自定义类型实现用途明确的枚举类型。但这并不能将取值范围限定在预定义的枚举值内。
	比如说，你预定义的枚举类型是byte，但是呢你传入了一个int类型的值，这是错误的。
	 */
	enum6()
}

func enum1() {

	/**
	iota是golang语言的常量计数器,只能在常量的表达式中使用;
	iota在const关键字出现时将被重置为0(const内部的第一行之前)，
	const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
	 */
	const (
		x = iota//0
		y        //1
		z        //2
	)

	println(x, y, z)//0 1 2
}

func enum2() {

	const (
		_ = iota        //0
		KB = 1 << (10 * iota)        //1 << (10 * 1) = 1 * 2^10 = 1 * 1024
		MB                        //1 << (10 * 2) = 1 * 2^20 = 1 * 1048576 = 1 * 1024 * 1024
		GB                        //1 << (10 * 3) = 1 * 2^30 = 1 * 1073741824 = 1 * 1024 * 1024 * 1024
	)

	println(KB, MB, GB)//1024 1048576 1073741824

}

/**
自增作用范围为常量组。可在多常量定义中使用多个iota，它们各自单独计数，只须确保组中每行常量的列数量相同即可
 */
func enum3() {
	const (
		_, _ = iota, iota * 10        //0 , 0 * 10
		a, b        //1,10
		c, d        //2,20
	)

	println(a, b, c, d)//1 10 2 20
}

/**
如中断 iota 自增，则必须显式恢复。且后续自增值按行序递增，而非 C enum 那般按上一取值递增
 */
func enum4() {
	const (
		a = iota        //0：第零行
		b                //1：第一行
		c = 100                //100：第二行
		d                //100(与上一行常量右值表达式相同)：第三行
		e = iota        //4（恢复 iota 自增，计数包括c、d，按行序递增）：第四行
		f                //5：第五行
	)

	println(a, b, c, d, e, f)//0 1 100 100 4 5
}

/**
自增默认类型为 int，可显式指定类型
 */
func enum5() {
	const (
		a = iota        //int
		b float32 = iota//float32
		c = iota        //int（如不显式指定 iota ，则与 b 数据类型相同）
	)

	println(a, b, c)//0 +1.000000e+000 2
}

type color byte

const (
	black color = iota        //0，指定常量类型
	red                        //1
	blue                        //2
)

/**
在实际编码中，建议用自定义类型实现用途明确的枚举类型。但这并不能将取值范围限定在预定义的枚举值内。
比如说，你预定义的枚举类型是byte，但是呢你传入了一个int类型的值，这是错误的。
*/
func enum6() {
	test(red)//1
	test(100)//100并未超出 color/byte 类型取值范围
	//x:=2
	//test(x)//错误：cannot use x (type int) as type color in argument to test
}

func test(c color) {
	println(c)
}
