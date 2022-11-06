package main

import (
	"fmt"
)

// Golang语言里的数组长度有限
// 并且顺序也是有序的
func main() {
	arr1 := [3]string{"你好", "再见"}
	fmt.Printf("%v\n%v\n", arr1[0], arr1[1])

	// 一个数组元素未满的情况下
	// 其他空元素则会被赋予与之对应的类型 "零" 值
	fmt.Println(arr1[2] == "")

	// 声明数组时 方括号内的 ... 表示
	// 根据给的字面量来确定数组的长度
	arr2 := [...]int{1, 23, 12}

	fmt.Println(arr2)

	// 遍历数组
	// 方式一:普通循环
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	// 方式二:range关键字
	for k, v := range arr2 {
		fmt.Println(k, v)
	}

	// 数组复制 -> 无论哪种情况
	// 参数传递、赋予新变量都是其副本 而非数组本身
	// 所以一般情况下别使用 值传递 而是使用 引用或指针传递
	// 或者是传递slice(切片)，好处是可以传递不同长度的数组
}
