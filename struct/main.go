package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// struct的变量创建语法
	// type 变量名 struct{
	// fieldName fieldType
	// fieldName fieldType
	// ...
	// }
	var locations struct {
		x int
		y int
	}
	locations.x = 21
	locations.y = 40

	fmt.Println(locations.x, locations.y)
	fmt.Println(locations)

	main2()
}

// 以上这种写法虽然可以创建一个struct实例
// 但是并不能很好的复用它 因为它只作用于局部

// 所以我们一般会创建一个struct类型来复用
// struct类型创建语法
// type 类型名 struct{
// fieldName fieldType
// fieldName fieldType
// ...
// }
type location struct {
	x int
	y int
	// 如果类型一样也可以这样定义字段名
	// x,y int
}

func main2() {
	// 使用自定义struct类型的复用
	var lc1 location
	lc1.x = 100
	lc1.y = 200
	var lc2 location
	lc2.x = 120
	lc2.y = 210

	// 通过复合字面值初始化struct
	// ! 遗漏的为赋值的变量编译器会赋予对应类型的0值
	// 方式一: 可以添加上字段名
	lc3 := location{x: 58, y: 80}
	// 方式二: 不添加字段名,按结构题字段名位置顺序分辨
	lc4 := location{67, 14}
	fmt.Println(lc1, lc2, lc3, lc4)

	// fmt.Printf的两种打印struct的方法
	// 方式一: 打印其字段名
	fmt.Printf("%+v\n", lc1)
	// 方式二: 不打印其字段名
	fmt.Printf("%v\n", lc1)

	main3()
}

// struct的复制
func main3() {
	// struct的复制并不像map类型一样
	// 操作复制变量并不会影响被复制原有的内容
	// 以下变量其实拥有各自独立的变量空间
	location1 := location{x: 12, y: 20}
	location2 := location1

	location2.x++
	location2.y++

	// ! 在函数传递struct当参数时也是值传递
	// 并不是和map一样是引用传递(指针传递)
	fmt.Printf("location1 = %+v\n", location1)
	fmt.Printf("location2 = %+v\n", location2)

	main4()
}

// 由struct组成的slice
func main4() {
	location1 := location{x: 10, y: 20}
	location2 := location{x: 30, y: 40}

	locations1 := []location{
		// 已有的struct变量不能与字面量初始化一起混用组成slice
		// location1,location2,location{x,50,y:60},
		location1, location2,
	}

	// 只允许全部使用字面量初始化struct slice
	// 或者全部使用struct初始化struct slice
	locations2 := []location{
		{x: 12, y: 210},
		{x: 100, y: 210},
		{x: 127, y: 640},
	}

	fmt.Printf("location1 slice = %+v\n", locations1)
	fmt.Printf("location2 slice = %+v\n", locations2)

	main5()
}

// 将struct编码为JSON格式
type user struct {
	Name, Email string
}

// 制作一个遇到错误就终止程序的函数
func eixtOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main5() {
	zxk := user{Name: "dsxksss", Email: "123456789@qq.com"}
	// encoding/json包中的Marshal函数可以将一个struct编码为JSON格式
	// ! Marshal函数只会对被导出(公开访问属性的值)的值编码
	// ! 也就是字段名首字母大写
	// Marshal函数返回两个值
	// 第一个是[]byte类型的silce
	// 第二个是err类型的错误
	bytes, error := json.Marshal(zxk)
	eixtOnError(error)

	// 因为Marshal类型返回的是[]byte类型的内容
	// 所以在打印其内容时要转换为string类型
	fmt.Println("zxk", string(bytes))

	// Golang语言json包中的函数都需要我们的struct字段名为大写才可以使用
	// 但有时候我们业务需求可能会使用开头字母小写的命名规范
	// 那么我们可以给字段注明一个 "标签",使得json包中的函数在运行时
	// 可以按标签里的样式去修改struct里对应的json字段名
	type tagUser struct {
		// ! 标签是使用``原始字符串标明的
		Name  string `json:"user_name"`
		Email string `json:"user_email"`
	}

	tagZxk := tagUser{Name: "dsxksss", Email: "123456789@qq.com"}
	tagBytes, tagError := json.Marshal(tagZxk)
	eixtOnError(tagError)

	fmt.Println("tagZxk", string(tagBytes))

	main6()
}

// Golang语言里没有类、对象、继承等面向对象里的东西
// 不过Golang语言提供struct、method
// 那我们就可以它们来获得类似的操作

// 将方法关联到struct
func (x user) showName() {
	fmt.Println(x.Name, "使用了user struct中的showName关联方法")
}

// 如果你想给struct给予更加复制的初始化内容
// 那么你可以制作一个构造函数来完成该操作
// ! Golang中没有专门用于构造用的函数
// ! 通常都是写一个普通函数,函数名以new或New开头
func newUser(x, y user) user {
	return user{x.Name, y.Email}
}

func main6() {
	dsxk := user{Name: "test", Email: "xxx@example.com"}
	dsxk.showName()
}

// ! 小知识
// ! 有一些用于构造的函数的名称就叫New
// ! 例如errors package 中的New函数
// ! 按理来说这样做的话函数名会重复
// ! 但是因为Golang中的函数调用时使用的是 包名.函数名 的形式
// ! 如果该函数叫做NewError 那么使用errors.NewError()调用
// ! 没有使用errors.New()简洁
