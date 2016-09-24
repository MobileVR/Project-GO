package main

import (
	"fmt"
	"os"
	"log"
)

/*
defer 误用
千万记住，延迟调用在函数结束时才被执行。
不合理的使用方式会浪费更多资源，甚至造成逻辑错误。
 */
func main() {

	/*
	案例：循环处理多个日志文件，不恰当的 defer 导致文件关闭时间延长。
	 */
	for i := 0; i < 10000; i++ {
		path := fmt.Sprintf("./log/%d.txt", i)

		f, err := os.Open(path)

		if err != nil {
			log.Println(err)
			continue
		}

		//这个关闭操作在 main 函数结束时才会执行，而不是当前循环中执行
		//这无端延长了逻辑结束时间和 f 的生命周期，平白多消耗了内存等资源
		defer f.Close()

	}

	/*
	应该直接调用，或重构为函数，将循环和处理算法分离。
	 */
	do := func(n int) {
		path := fmt.Sprintf("./log/%d.txt", n)

		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			continue
		}

		//延迟调用用在此匿名函数结束时执行，而非 main
		defer f.Close()
	}

	for i := 0; i < 10000; i++ {
		do(i)
	}

}
