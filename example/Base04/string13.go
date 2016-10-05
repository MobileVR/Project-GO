package main

import (
	"unicode/utf8"
)

/*
类型 rune 专门用来存储 Unicode 码点（code point），它是 int32 的别名，相当于 UCS-4/UTF-32 编码格式。
使用单引号的字面量，其默认类型就是 rune。
 */
/*func main() {

	f := '我'

	fmt.Printf("%T\n", f)//int32

}*/

/*
除 []rune 外，还可直接在 rune、byte、string 间进行转换。
 */
/*func main() {

	r := '我'

	s := string(r)//rune to string

	fmt.Println(s)//我

	b := byte(r)//rune to byte

	fmt.Println(b)//17

	s2 := string(b)//byte to string

	fmt.Println(s2)//

	r2 := rune(b)//byte to rune

	fmt.Println(r2)//17

}*/

/*
要知道字符串存储的字节数组，不一定就是合法的 UTF-8 文本。
 */
/*func main() {

	s := "你好"

	s = string(s[0:1] + s[3:4])//截取并拼接一个“不合法”的字符串

	fmt.Println(s, utf8.ValidString(s))//�� false
}*/

/*
标准库 unicode 里提供了丰富的操作函数。除验证函数外，还可用 RuneCountInString 代替 len 返回准确的 Unicode 字符数量。
 */
func main() {

	s := "你.好"

	println(len(s), utf8.RuneCountInString(s))//7 3

}
