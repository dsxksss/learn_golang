package main

import "fmt"

// 在面向对象的世界中 对象是由更小的对象组合而成的
// 术语: 对象组合 或者 组合
// Golang通过struct实现组合(composition)
// Golang提供了"嵌入"(embedding)特性 它可以实现方法的转发
// 组合是一种更简单、灵活的方式 面向对象的内容都可以通过组合实现

// 以下这个struct可以优化成
// 多个struct组合后的struct
// type report struct {
// 	sol       int
// 	high, low float64
// 	lat, long float64
// }

// 例子
type report struct {
	sol int
	// 如果字段名和字段类型一致
	// 那么可以省略只写一项
	temperature
	location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type celsius float64

//

func main() {
	bradbury := location{-4.4154, 137.4671}
	t := temperature{high: -1.0, low: -78.0}
	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.temperature.high)
}
