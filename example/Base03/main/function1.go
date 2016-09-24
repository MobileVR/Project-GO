package main

/*
1、无须前置声明
2、不支持命名嵌套定义（nested）
3、不支持同名函数重载（overload）
4、不支持默认参数
5、支持不定长参数
6、支持多返回值
7、支持命名返回值
8、支持匿名函数和闭包
 */
func main() {

	/*
	不支持命名嵌套定义
	 */
	/*func add(x,y int) int {//syntax error: unexpected add, expecting (
		return x + y
	}*/

}

/**
左花括号不能另起一行
 */
/*func test()
{//syntax error: unexpected semicolon or newline before {
}*/

func test(){
}

/*
不支持同名函数重载
 */
func test(x int) {//test redeclared in this block
}
