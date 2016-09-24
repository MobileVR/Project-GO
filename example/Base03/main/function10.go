package main

import (
	"time"
	"log"
)

/*
如果函数参数过多，建议将其重构为一个复合结构类型，也算是变相实现可选参数和命名实参功能。
将过多的参数独立成 option struct，既便于扩展参数集，也方便通过 newOption 函数设置默认配置。
这也是代码复用的一种方式，避免多处调用时烦琐的参数配置。
 */
func main() {

	opt := newOption()

	opt.port = 8085//命名参数设置

	server(opt)

}

type serverOption struct {
	address string
	port    int
	path    string
	timeOut time.Duration
	log     *log.Logger
}

func newOption() *serverOption {
	return &serverOption{//默认参数
		address:"0.0.0.0",
		port:8080,
		path:"/var/test",
		timeOut:time.Second * 5,
		log:nil,
	}
}

func server(option *serverOption) {
	println(option.port)
}

