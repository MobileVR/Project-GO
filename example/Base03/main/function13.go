package main

import (
	"errors"
	"fmt"
)

/*
稍有不便的是没有元组（tuple）类型【元组(tuple) 把多个值组合成一个复合值。元组内的值可以使任意类型,并不要求是相同类型】，
也不能用数组、切片接收，但可用 "_" 忽略掉不想要的返回值。
多返回值可用作其他函数调用实参，或当作结果直接返回。
 */
func main() {

	//多返回值用作实参
	log13(test13())//0 division by zero

}

func div13(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("division by zero")
	}

	return x / y, nil

}

func log13(x int, err error) {
	fmt.Println(x, err)
}

func test13() (int, error) {
	return div13(5, 0)//多返回值用作 return 结果
}