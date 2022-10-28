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

	for i := 0; i < 5; i++ {
		fmt.Printf("i=%v\n", i)
	}

	shalom := "shalom"
	// 这里的len函数返回的是该字符串的bytes
	// 如果要获取字符串的长度的话要导入unicode/utf8包
	// 使用utf8.RuneCountInString函数来获取长度
	for i := 0; i < len(shalom); i++ {
		fmt.Printf("%c\n", shalom[i])
	}

	// 要遍历某个类型的话可以使用range关键字
	// 该关键字是一个迭代器的语法糖,可以遍历各种集合
	for i, c := range shalom {
		fmt.Printf("i=%v c=%c\n", i, c)
	}

	fmt.Println("循环结束")
}
