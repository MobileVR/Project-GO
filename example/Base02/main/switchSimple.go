package main

func main() {

	/*
	switch基本格式
	 */
	//switch1()

	/*
	switch同样支持初始化语句，按从上到下、从左到右顺序匹配 case 执行。只有全部匹配失败时，才会执行 default 块。
	考虑到 default 作用类似于 else ，建议将其放置在 switch 末尾。
	 */
	//switch2()

	/*
	相邻的空 case 不构成多条件匹配
	 */
	//switch3()

	/*
	不能出现重复的 case 常量值
	 */
	//switch4()

	/*
	无须显式执行 break 语句，case 执行完毕后自动中断。
	如须贯通后续 case （源码顺序），须执行 fallthrough ，但不再匹配后续条件表达式。
	 */
	//switch5()

	/*
	注意：fallthrough 必须放在 case 块结尾，可使用 break 语句阻止。
	 */
	//switch6()

	/*
	某些时候，switch 还被用来替换 if 语句。被省略的 switch 条件表达式默认值为 true，继而与 case 比较表达式结果匹配。
	 */
	//switch7()

}

/*
switch基本格式
 */
func switch1() {

	a, b, c, x := 1, 2, 3, 2

	switch x {//将 x 与 case 条件匹配
	case a:
		println("a")
	case b, c://多个匹配条件命中其中一条即可（OR），变量
		println("b,c")
	case 4://常量
		println("4")
	default:
		println("default")
	}

}

/*
switch同样支持初始化语句，按从上到下、从左到右顺序匹配 case 执行。只有全部匹配失败时，才会执行 default 块。
考虑到 default 作用类似于 else ，建议将其放置在 switch 末尾。
 */
func switch2() {

	switch q := 5; q {
	default://编译器确保不会先执行 default 块
		println("default")
	case 5:
		println("q")
	}

}

/*
相邻的空 case 不构成多条件匹配
 */
func switch3() {

	x := 1

	switch x {
	case 1://单条件，内容为空。隐式"case 1:break;"
	case 2:
		println("2")
	case 3:
		println("3")
	default:
		println("default")
	}

}

/*
不能出现重复的 case 常量值
 */
func switch4() {
	switch x := 2; x {
	case 1:
		println("1")
	case 2:
		println("2")
	//case 2://duplicate case 2 in switch
	//	println("3")
	}
}

/*
无须显式执行 break 语句，case 执行完毕后自动中断。
如须贯通后续 case （源码顺序），须执行 fallthrough ，但不再匹配后续条件表达式。
 */
func switch5() {
	switch x := 5; x {
	default:
		println("default")
	case 5:
		println("5")
		fallthrough//继续执行下一个case，但不再匹配条件表达式
	case 6:
		println("6")

	/*
	如果在此继续 fallthrough ，不会执行 default ，完全按源码顺序
	导致 'cannot fallthrough final case in switch' 错误
	 */
	//fallthrough//cannot fallthrough final case in switch

	}
}

/*
注意：fallthrough 必须放在 case 块结尾，可使用 break 语句阻止。
 */
func switch6() {
	switch x := 5; x {
	case 5:

		if true {
			println("break")
			break//终止，不再执行后续语句
		}

		println("5")

		fallthrough//必须是 case 块的最后一条语句

	case 6:
		println("6")
	}
}

/*
某些时候，switch 还被用来替换 if 语句。被省略的 switch 条件表达式默认值为 true，继而与 case 比较表达式结果匹配。
 */
func switch7() {
	switch x := 5; {//相当于"switch x := 5; true { ... }"
	case x > 5:
		println("x > 5")
	case x > 0&&x <= 5://不能写成"case x > 0,x <= 5"，因为多条件是 OR 关系
		println("x > 0 && x <= 5")
	default:
		println("default")
	}
}
