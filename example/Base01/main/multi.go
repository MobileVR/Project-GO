package main

/**
多赋值运算
 */
func main() {

	//多赋值运算
	multi1()
}

/**
多赋值运算
 */
func multi1() {

	x, y := 1, 2

	//先计算出右边的值y+3、x+2，然后再对x、y变量赋值
	x, y = y + 3, x + 2

	println(x, y)//5 3

}
