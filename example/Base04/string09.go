package main

import "fmt"

/*
用 append 函数，可将 string 直接追加到 []byte 内。
 */
func main() {

	var bs []byte

	bs = append(bs,"abc"...)

	fmt.Println(bs)//[97 98 99]

}
