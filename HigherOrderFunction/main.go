package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 高阶函数(一等函数)
// 在Go里 函数是头等的 它可以用在整数、字符串或其他类型能用的地方
// - 将函数赋给变量
// - 将函数作为参数传递给函数
// - 将函数作为函数的返回类型

// 例子1: 将函数赋给变量
type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

// 例子2: 将函数当作函数参数传入
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° k\n", k)
		time.Sleep(time.Second)
	}
}

// 例子3: 匿名函数
var f = func() {}

func main() {
	// 函数存入变量
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	// 函数当作参数
	measureTemperature(3, fakeSensor)

	// 使用匿名函数
	f()
}
