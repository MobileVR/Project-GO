package main

/*
要实现传出参数（out），通常建议使用返回值。当然，也可以继续用二级指针。
 */
func main() {

	var p *int

	println(p, &p)//0x0 0xc04202df18

	function9(&p)

	println(p, *p)//0xc042038000 100

}

func function9(p **int) {

	x := 100

	*p = &x

	println(p, *p)//0xc04202df18 0xc042038000

}
