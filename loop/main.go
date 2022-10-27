package main

import (
	"fmt"
	"time"
)

func main() {
	// for 后面如果没有跟条件,那就是无限循环
	// - 也和其他语言一样可以使用break跳出循环
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second) //阻塞一秒
		count--
	}
	fmt.Println("循环结束")
}
