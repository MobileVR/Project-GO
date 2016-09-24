package main

import (
	"fmt"
	"log"
)

/*
自定义错误
自定义错误类型通常以 Error 为名称后缀。
在用 switch 按类型匹配时，注意 case 顺序。
应将自定义类型放在前面，优先匹配更具体的错误类型。
大量函数和方法返回 error，使得调用代码变得很难看，一堆堆的检查语句充斥在代码行间。
解决思路有：
1.使用专门的检查函数处理错误逻辑（比如记录日志），简化检查代码。
2.在不影响逻辑的情况下，使用 defer 延后处理错误状态（err退化赋值）。
3.在不中断逻辑的情况下，将错误作为内部状态保存，等最终“提交”时再处理。
 */
func main() {

	z, err := div21(1, 0)

	if err != nil {

		switch e := err.(type) {//根据类型匹配
		case DivError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}

		log.Fatalln(err)

	}

	println(z)

}

type DivError struct {
	//自定义错误类型
	x, y int
}

func (DivError) Error() string {
	//实现 error 接口方法
	return "division by zero"
}

func div21(x, y int) (int, error) {

	if y == 0 {
		return 0, DivError{x, y}//返回自定义错误类型
	}

	return x / y, nil

}
