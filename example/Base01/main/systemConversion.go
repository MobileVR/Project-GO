package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
进制间转换
 */
func main() {

	//利用进制占位符输出
	//system1()

	//利用 strconv.ParseInt 方法进行进制间的转换
	//system2()

}

/**
利用进制占位符输出
 */
func system1() {

	a, b, c := 100, 0144, 0x64

	println(a, b, c)//100 100 100
	fmt.Printf("0b%b,%#o,%#x\n", a, a, a)//0b1100100,0144,0x64
	fmt.Println(math.MinInt8, math.MaxInt8)//-128 127

}

/**
利用 strconv.ParseInt 方法进行进制间的转换
 */
func system2() {

	a, _ := strconv.ParseInt("1100100", 2, 32);
	b, _ := strconv.ParseInt("0144", 8, 32);
	c, _ := strconv.ParseInt("64", 16, 32);

	println(a, b, c)//100 100 100

	println("0b" + strconv.FormatInt(a, 2))//0b1100100
	println("0" + strconv.FormatInt(b, 8))//0144
	println("0x" + strconv.FormatInt(c, 16))//0x64

}
