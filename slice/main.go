package main

import (
	"fmt"
	"sort"
)

// 因为参数传递、赋予新变量都是其副本 而非数组本身
// 所以一般情况下别使用 值传递 而是使用 引用或指针传递
// 或者是传递slice(切片)，好处是可以传递不同长度的数组

func main() {
	// 声明数组时 方括号内的 ... 表示
	// 根据给的字面量来确定数组的长度
	arr := [...]int{1, 23, 12}

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
	fmt.Println(arr[:2])
	fmt.Println(arr[0:])
	fmt.Println(arr[1:2])
	fmt.Println(arr[:])

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

	// append 函数
	// slice中的个数决定了slice的长度
	// 如果slice的底层数组比slice还大
	// 那么就说明该slice还有容量可供增长
	slice1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	slice1 = append(slice1, "Bob", "Jack", "Sary")
	fmt.Println("append", slice1)

	main3()
}

func dump(label string, slice []string) {
	// slice长度 -> 切片里有多少个元素 那么就是其长度
	// slice容量 -> 底层数组的长度
	// 获取长度用len函数
	// 获取容量用cap函数
	fmt.Printf("%v:length %v,capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main3() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// 这里的dwarfs1 存储的容量超过了自身的切片长度
	// 所以在添加了内容超过自身的容量之后就会对其底层数组增加一倍
	dwarfs2 := append(dwarfs1, "Bob1", "Jack1")
	dwarfs3 := append(dwarfs2, "Bob2", "Jack2", "Sary2")

	// 也就是这里的底层数组只出现了两个
	// 第一个作用于dwarfs1
	// 第二个作用于dwarfs2 和 dwarfs3
	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
	dump("dwarfs[1:2]", dwarfs1[1:2])

	// 因为dwarfs2 和 dwarfs3公用一个底层数组
	// 所以他们的内容其实是互相绑定的
	// 这里修改了dwarfs3里的内容
	dwarfs3[1] = "xxx"
	// 以下的dwarfs1 和 dwarfs2 其中的
	// dwarfs2才发生了变化 dwarfs1内容还是保持不变
	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)

	// slice切片的第三个参
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// [0:4:4] 里的最后一个4 表示
	// 限制了terrestrial的容量长度
	terrestrial := planets[0:4:4]
	// 当往里面添加元素是发现底层数组容量不够
	// 所以在赋值给worldes时又增加了一份容量进去
	worldes := append(terrestrial, "Ceres")

	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worldes", worldes)

	main4()
}

func main4() {
	// 利用make函数对slice进行预分配
	// 我们知道当slice的容量不足以执行append函数时
	// golang必须创建新的数组并复制数组中的内容
	// 但是可以通过make避免额外的内存分配和数组复制

	// 这里利用了make函数创建了一个长度为0容量为10的一个slice切片
	dwarfs := make([]string, 0, 10)

	// make 函数为两个参数时 第二个参数表示长度和容量
	// 当为三给参数时 第二个参数表示长度 第三个表示容量
	dump("before make(dwarfs)", dwarfs)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris", "xxx")

	dump("after make(dwarfs)", dwarfs)

}
