package main

import "fmt"

// 指针是指向另一个变量地址的变量
// Golang的指针同时也强调安全性 不会出现迷途指针

// &操作符用于取变量地址符 不能取字面量地址
// *操作符用于解引用 提供内存地址指向的值
func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println("解地址前", address)
	fmt.Println("解地址后", *address)
}
