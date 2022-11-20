package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//  处理错误
//  Golang语言允许函数和方法同时返回多个值
// 按照惯例 函数在返回错误时 最后边的返回值应该用来表示错误
// 调用函数后 应立即检测是否发生错误
// - 如果没有错误发生 那么返回的错误值应为nil

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		// ! 注意 当错误发生时 函数返回的其他值就不太可信了
		fmt.Println(err)
		os.Exit(1) // Exit函数只要传入非零值则表示该程序发生了错误
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	main2()
}

// 如何优雅的处理错误
// 减少错误处理代码的一种策略是
// 将程序中不会出错的部分和包含潜在
// 错误隐患的部分隔离开来
// 对于不得不返回错误的代码 应尽力简化相应的错误处理代码

// 文件处理

func proverbs(name string) error {
	// create -> 创建文件
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// defer关键字
	// 使用defer关键字 Golang
	// 可以确保所有deferred的动作可以在函数返回前执行
	// defer虽然并不是专门做错误处理的
	// 但defer可以消除必须时刻惦记释放执行资源的负担
	// 这里这样做的好处是 不管在何时return前 都会执行defer的动作
	defer f.Close()

	// Fprintln -> 写入数据
	_, err = fmt.Fprintln(f, "Error are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err
}

func main2() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	main3()
}

// 有创意的错误处理
// 例子:
type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	// 这种写法就可以避免重复的
	// 返回错误了
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

// 利用结构体方法后的一种
// 创新的错误处理方式
func newProverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errprs are values1")
	sw.writeln("Errprs are values2")
	sw.writeln("Errprs are values3")
	sw.writeln("Errprs are values4")
	sw.writeln("Errprs are values5")
	return sw.err
}

// 生成错误
func createError() error {
	// 错误信息应具有信息性
	// 可以把错误信息当作用户界面的一部分
	// 无论对最终用户还是开发者

	// 按照惯例
	// 包含错误信息的变量名
	// 应以Err开头
	// errors.New(错误信息)
	ErrCreateError := errors.New("发生了错误")

	// errors.New()函数是使用指针实现的
	// 所以你不能用来比较两个err的错误文字是否相同

	return ErrCreateError
}

// 自定义错误类型
// error 类型其实是一个内置的接口
// 任何类型只要实现了返回string的Error()方法
// 则就满足了该接口
// 我们则可以基于这个接口来创建新的错误类型

// 按照惯例 自定义错误类型的名字应以Error结尾
// - 有时候名字就是Error 例如 url.Error
type newError []error

func (ne newError) Error() string {
	var s []string
	for _, err := range ne {
		s = append(s, err.Error())
	}
	return strings.Join(s, "-")
}

func main3() {
	Err1 := errors.New("1")
	Err2 := errors.New("2")
	Err3 := errors.New("3")
	Err4 := errors.New("4")
	Err5 := errors.New("5")
	Errs := newError{Err1, Err2, Err3, Err4, Err5}
	fmt.Println(Errs.Error())

	// 类型断言
	// 我们可以使用类型断言来访问每一种错误
	// 使用类型断言 你可以把接口类型转换成底层的具体类型
	// - 例如 err.(newError)
	// 如果类型满足多个接口 那么类型断言可使它从一个接口类型
	// 转化为另一个接口类型
	if errs, ok := Errs[0].(newError); ok {
		for _, e := range errs {
			fmt.Println("类型匹配成功", e)
		}
	} else {
		fmt.Println("该类型不是newError类型")
	}

	main4()
}

// Golang没有异常 它有个类似的机制panic
// 当panic发生 那么程序就会奔溃

// 其他语言里的异常 vs Golang的错误值
// 其他语言的异常在行为和实现上与Golang语言的错误值有很大的不同
// - 如果函数抛出异常 并且附近没人捕获它 那么它就会"冒泡"到函数的调用者那里
// - - 如果还没有人进行捕获 那么就继续"冒泡"到更上层的调用者那里
// - - 直到达到栈(Stack)的顶部 例如main函数
// - 异常这种错误处理方式可被看作是可选的:
// -- 不处理异常 就不需要加入其他代码
// -- 想要处理异常 就需要加入相当数量的专用代码

//	Golang语言中的错误值更简单灵活
// - 忽略错误是有意识的决定 从代码上看也是显而易见的
// 实际上panic很少出现
// 使用panic()函数创建panic

func main4() {
	// 通常更推荐使用错误值 其次才是panic
	// panic 比 os.Exit更好 panic 后会执行所有的defer的动作而os.Exit则不会
	// 有时候Golang程序会panic而不是返回错误值(例如除以0这种情况)

	// 为了防止panic导致程序奔溃
	// Golang提供了recover函数
	// defer的动作会在函数返回前执行 即使发生了panic
	// 但如果defer的函数调用了recover  panic就会停止 程序将继续运行
	// 例子
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("出现panic")
}
