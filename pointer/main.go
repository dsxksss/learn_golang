package main

import (
	"fmt"
	"strings"
)

type person struct {
	name, superpower string
	age              int
}

// 指针是指向另一个变量地址的变量
// Golang的指针同时也强调安全性 不会出现迷途指针

// &操作符用于取变量地址符 不能取字面量地址
// *操作符用于解引用 提供内存地址指向的值
func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println("解地址前", address)
	fmt.Println("解地址后", *address)

	canada := "Canada"
	// 将*放在类型前面表示声明指针类型
	// 将*放在变量前面表示解引用操作
	var home *string
	fmt.Printf("home is a %T\n", home)
	home = &canada
	fmt.Println(*home)

	// canada内容发生变化
	// 指向它的指针内容也会跟随它变化
	canada = "Canadaxxx"
	fmt.Println(*home)

	// 指针变量也可以操纵所指向的地址中的内容
	*home = "jack"
	fmt.Println(canada)

	// 新指针变量也可以操纵给予给它的
	// 另一个指针变量地址中的内容
	newHome := home
	*newHome = "xxx"
	fmt.Println(canada)

	// 如果两个指针变量持有相同的内存地址
	// 那么它们就是相等的
	fmt.Println(home == newHome)

	// 如果将指针指向地址中的内容赋值给其他变量
	// 那么他们两是互不影响的 因为只会复制一份给对方
	light := *home
	*home = "abcd"
	fmt.Println(light)
	fmt.Println(canada)

	// 指向结构的指针
	// 与字符串和数值不一样 复合字面量的前面可以放置&

	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	// 如果是指向的复合字面量
	// 那么解引用符号可以不加使用
	// 或者使用强转使用(*timmy)
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	// 指向数组的指针
	// 和结构体一样 可以吧&放在数组的
	// 复合字面值前面来创建指向数组的指针
	// 数组值执行索引或切片操作时会自动解引用
	// 没有必要写成(*array)[0]这种形式
	// 与C语言不一样 Golang里面数组和指针是两种完全独立的类型
	// Slice和map的复合字面值也可以这前面放置&操作赋
	// 但是Golang并没有为它们提供自动解引用的功能
	array := &[3]string{
		"flight",
		"invisibility",
		"super strength",
	}
	fmt.Println(array[0])
	fmt.Println(array[1:2])

	main2()
}

// Golang语言的函数和方法都是按值传递参数的
// 这意味着函数总是操作于被传递参数的副本

func birthday(p *person) {
	p.age++
}

// 但指针被传递到函数时 函数将接收传入的内存地址的副本
// 之后函数可以通过解引用内存地址来修改指针指向的值
func main2() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)

	main3()
}

// 指针接收者
// 方法的接收者和方法的参数值处理指针方面是很相似的

func (p *person) birthday() {
	p.age++
}

func main3() {
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)

	nathan := person{
		name: "Nathan",
		age:  17,
	}
	// 虽然接收者并不是一个指针类型
	// 但是在我们使用.操作符的时候 Golang
	// 会自动使用&取得变量的内存地址
	nathan.birthday()
	// 所以也可以不使用下面这种方法执行
	// (&nathan).birthday()
	fmt.Printf("%+v\n", nathan)

	// ! 使用指针作为接收者的策略应该始终如一
	// ! 如果一种类型的某些方法需要使用到指针作为接收者
	// ! 就应该为这种类型的所有方法都使用指针作为接收者

	main4()
}

// 内部指针
// Golang语言提供了内部指针这种特性
// 它用于确定struct中的制定字段的内存地址

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func main4() {
	player := character{name: "Matthias"}
	// &操作符不仅可以获得struct的内存地址
	// 还可以获得struct中指定字段的内存地址
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)

	main5()
}

// 修改数组
// 函数通过指针对数组的元素进行修改

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func main5() {
	var board [8][8]rune
	fmt.Printf("函数通过指针修改数组元素前%c\n", board[0][0])
	reset(&board)
	fmt.Printf("函数通过指针修改数组元素后%c\n", board[0][0])

	// 隐式的指针
	// Golang语言里的一些内置的集合类型
	// 就在暗中使用指针:
	// - map 在被赋值或者被作为参数传递的时候不会被复制(详细见map文件)
	// - - map 就是一种隐式指针
	// - - 这种写法就是多此一举:
	// func demolish(planets *map[string]string)
	// - map 的KEY和VALUE都可以是指针类型
	// 需要将指针指向map的情况并不多见

	// - slice slice是指向数组的窗口
	// - - 实际上slice在指向数组元素时也使用了指针
	// 每个slice内部都会被表示为一个包含了3个元素的结构
	// 它们分别指向:
	// - 数组的指针
	// - slice的容量
	// - slice的长度
	// 当slice被直接传递至函数或方法时
	// slice的内部指针及可以对底层数据进行修改

	main6()
}

// 指向slice的显式指针的唯一作用就是
// 修改slice本身 slice 的长度、容量、起始偏移量

func reclassify(planets *[]string) {
	// 这里修改了slice自身的长度
	*planets = (*planets)[0:8]
}

func main6() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	// 原先planets有九个元素
	fmt.Println("reclassify 修改前", planets)
	reclassify(&planets)
	fmt.Println("reclassify 修改后", planets)

	main7()
}

// 指针和接口
type talker interface {
	talk() string
}

func shout(t talker) {
	fmt.Println(strings.ToUpper(t.talk()))
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

// 如果方法使用的是指针接受者 那么情况会有所不同
type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main7() {
	// 本例中 无论martian还是指向martian的指针
	// 都可以满足talker接口,自然shout函数也可以接受了
	shout(martian{})
	shout(&martian{})

	pew := laser(2)
	shout(&pew)

	// ! 应合理使用指针 不要过度使用指针
}
