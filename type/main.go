package main

import (
	"fmt"
	"strconv"
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

	// 类型转换
	// 例子 string(b) uint8(c)
	// string只会转换为对应的codepoint
	// 如果要将一个数字转换为string类型
	// 则要使用strconv.Itoa 函数来转换
	str := "abcde" + strconv.Itoa(int(b)) + "secx"
	fmt.Println(str)

	// 当然日常使用最多的字符串强转类型是使用
	// fmt.Sprintf函数,返回值为一个string类型
	countdown := 9
	str1 := fmt.Sprintf("今年你%v岁了", countdown)
	fmt.Println(str1)

	// nil 关键字
	// 如果err返回的内容为nil则表示没有发生错误
	countdown2, err := strconv.Atoi("10")
	if err != nil {
		// 发生错误时
		fmt.Println(countdown2)
	}
	fmt.Println(countdown2)

	// 使用type关键字来创建新类型
	type newfloat float64
	var x newfloat = 21.2
	fmt.Println(x)

}
