package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		a = 10
		b = 20
	)
	fmt.Println(a)
	fmt.Println(b)
	var num = rand.Intn(20) + 1
	fmt.Println(num)
	fmt.Println("你好")
}
