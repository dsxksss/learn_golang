package main

import (
	"fmt"
)

// golang基本类型
// 整型 - 10种 5种有符号类型 5种无符号类型(u前缀) 前8种与CPU架构无关系
// - int8                             -128 ～ 127
// - uint8                               0 ～ 255
// - int16                         -32,768 ～ 32,767
// - uint16                              0 ～ 65,535
// - int32                  -2,147,483,648 ～ 2,147,483,647
// - uint32                              0 ～ 4,294,967,295
// - int64      -9,223,372,036,854,775,808 ～ 9,223,372,036,854,775,807
// - uint64                              0 ～ 18,446,744,073,709,551,615
// - int  -> 针对架构 老移动设备 int类型表示 int32 新设备 int类型表示 int64
// - uint -> 针对架构 老移动设备 uint类型表示 uint32 新设备 int类型表示 uint64

func main() {
	// 在Printf里可以使用%T输出类型
	var a uint = 3212
	b := 210.2
	var c = 43
	fmt.Printf("%T\n", a) // 显示 uint
	fmt.Printf("%T\n", b) // 显示 float64
	fmt.Printf("%T\n", c) // 显示 int
}
