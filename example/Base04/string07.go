package main

import "fmt"

/*
使用 for 遍历字符串时，分 byte 和 rune 两种方式，两者是有区别的。
最直观的区别就是
rune 能操作任何字符
byte 不支持中文的操作
 */
func main() {

	//s := "abcdefg"

	/*
	输出结果示例：
	0:[a]
	1:[b]
	2:[c]
	3:[d]
	4:[e]
	5:[f]
	6:[g]
	 */
	/*for i := 0; i < len(s); i++ {
		//byte
		fmt.Printf("%d:[%c]\n", i, s[i])
	}*/

	s1 := "我们"

	/*
	输出结果示例：
	0:[æ]
	1:[]
	2:[]
	3:[ä]
	4:[»]
	5:[¬]
	 */
	for i := 0; i < len(s1); i++ {
		//byte
		fmt.Printf("%d:[%c]\n", i, s1[i])
	}

	/*
	输出结果示例：
	0,[我]
	3,[们]
	 */
	for i, c := range s1 {//rune:返回数组索引号，以及 Unicode 字符
		fmt.Printf("%d,[%c]\n", i, c)
	}

}
