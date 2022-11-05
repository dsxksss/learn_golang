package main

import (
	"fmt"
	"sort"
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
	// Golang里面很多函数都倾向于使用slice而不是数组作为参数
	// 想要获得于底层数组相同元素的slice 那么可以使用[:]进行切分
	// 切分数组并不是创建slice的唯一方法 可以之间声明slice:
	// - 例如[]string []内不写任何内容即可声明slice

	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfSlice := dwarfArray[:]

	dwarfS := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Println(dwarfSlice, dwarfS)
	fmt.Printf("array Type=%T , slice Type=%T", dwarfArray, dwarfS)

	// 数组的切片
	// Golang数组切片不能用负数
	// 小提示:切片语法也可以使用在
	// Golang的字符串中
	// 只不过切分字符串中的索引值代表的是
	// 字节数而非rune(字符)个数
	fmt.Println(arr2[:2])
	fmt.Println(arr2[0:])
	fmt.Println(arr2[1:2])
	fmt.Println(arr2[:])

	// 带方法的slice
	main2()
}

type StringSlice []string // 其实sort包内自带这个类型,去掉这行也能执行

func (p StringSlice) Sort() { // sort包内当然也有Sort函数,所以这行去掉也能执行
}

func main2() {
	planets := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println("before", planets)
	sort.StringSlice(planets).Sort()
	fmt.Println("after", planets)
}
