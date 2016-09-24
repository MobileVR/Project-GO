package main

/*
函数属于第一类对象，具备相同签名（参数及返回值列表）的视作同一类型
 */
func main() {

	f := hello
	//exec(f())//f() used as value
	exec(f)//hello word

}

func hello() {
	println("hello word")
}

func exec(f func()) {
	f()
}
