package main

import "errors"

// 如果两个参数类型一样可以指表明后面参数的类型
// 函数名大写表示public 小写表示private
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 函数可以返回多个返回值
// func maxE(a, b int) (result int, err error) {
func maxE(a, b int) (int, error) {
	if a > b {
		// nil 表示没有错误
		return a, nil
	} else if a < b {
		return b, nil
	}
	return -1, errors.New("发生错误")
}

// 可变参数函数
// ...表示函数参数的数量是可变的
// interface{} 表示空接口 类似于ts中的any
func printx(a ...interface{}) (interface{}, error) {
	return a, errors.New("我是一个可变参数函数")
}
