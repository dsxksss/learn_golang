package main

import (
	"fmt"
)

func main() {
	var age = 18
	if age > 18 {
		fmt.Println("成年")
	} else if age == 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	if age > 18 && age == 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	var a = true
	var b = false

	if !a || !b {
		fmt.Println("我输出了")
	}

	var swq = "abc"

	switch swq {
	case "abc":
		fmt.Println("正确")
		fallthrough // fallthrough会不判断执行下一个case
	case "xxx":
		fmt.Println("错误")
		fallthrough
	default:
		fmt.Println("默认")
	}
}
