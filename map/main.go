package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// 声明语法
	// 方式1、var map1 = map[string]int{k,v}
	// 方式2、map1 := map[string]int{k:v}
	// []里的类型是key的类型
	// []外的类型是value的类型
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	// 通过在[]里填写key值获取对应的value
	fmt.Println(temperature["Earth"])

	// 修改map里key对应的value内容
	temperature["Earth"] = 20
	// 往map里增加新内容
	temperature["Venus"] = 478
	fmt.Println(temperature)

	// 如果不存在该key|value的话则返回对应的0值
	fmt.Println(temperature["Jack"])

	// 如果有判断是否存在该key|value的业务逻辑时
	// 我们可以使用,ok写法 -> {ok名字不是固定的 可自定义}
	// 它会返回两个值 一个是通过key找到的value值(如果有的话)
	// 另一个是一个bool值 如果有的话则为true 否则false
	if jack, ok := temperature["Jack"]; ok {
		fmt.Println("存在", jack)
	} else {
		fmt.Println("没找到")
	}

	// map的传递细节
	// map和其他类型不一样 它的不会被复制 不会像数组一样创建相应的副本
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets
	planets["Mars"] = "whoops"
	fmt.Println("planets", planets)
	fmt.Println("planetsMarkII", planetsMarkII)

	// delete 函数
	// 可以删除map内的内容
	// - p1 : 要删除谁里的内容
	// - p2 : 要删除具体哪个
	delete(planets, "Earth")
	fmt.Println("after delete planets", planets)
	fmt.Println("after delete planetsMarkII", planetsMarkII)

	// 使用make函数对map进行预分配
	// 除非你使用复合字面值来初始化map  像这样 -> x := map[string]int{"你好",20}
	// 否则必须使用内置的make函数来为map分配空间
	// 创建map时 make函数可以接受一个或者两个参数
	// - p1 : 要创建的数据类型
	// - p2 : 用于为指定数量的key预先分配空间
	eMap := make(map[float64]int, 8)
	// 注! 使用make函数创建的map的初始长度为0
	fmt.Println(len(eMap))

	// 使用map制作计数器
	// 功能: 记录了出现温度的次数
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	freauency := make(map[float64]int)
	// _ -> index
	// t -> item
	for _, t := range temperatures {
		freauency[t]++
	}
	// t    -> key
	// num  -> value
	// 使用range遍历map的时候顺序是无法保证的
	for t, num := range freauency {
		fmt.Printf("%+.2f occurs %d time\n", t, num)
	}

	// 使用slice和map实现数据分组
	// 此处还会使用上面的temperatures slice
	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		// Trunc 函数是截断一个float类型为整数的功能
		// 例子: -28.0/10 = -2.8
		g := math.Trunc(t/10) * 10
		// 按温度的跨度来进行分组
		// 例如:
		// -28.0、-21.4、21.0 为一组
		// -38.4、34.5、-31.45 为一组
		groups[g] = append(groups[g], t)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}

	// 将map用作set
	// - set这种集合类似于数组 但是元素不会重复
	// - 但是Golang语言里没有提供set类型
	// 此处还会使用上面的temperatures slice
	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}
	if _, ok := set[-28.0]; ok {
		fmt.Println("set member")
	} else {
		fmt.Println("not find")
	}
	fmt.Println(set)

	// 因为输出顺序不能保证
	// 这里使用一些小手段来按顺序输出
	// 先创建一个切片
	unique := make([]float64, 0, len(set))
	for t := range set {
		// 循环给刚刚创建的unique切片赋值
		unique = append(unique, t)
	}
	// Float64s函数 按从小到大顺序排列
	// 会改变原切片内容
	sort.Float64s(unique)
	fmt.Println(unique)
}
