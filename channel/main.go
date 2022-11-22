package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 需要使用多种线程或者并发的去执行任务时
// 可以使用goroutine
// 这个东西类似于其他编程语言里的线程或者进程等
// 只需要在执行的任务面前加上go关键字
// 即可并发的执行该内容
// 但是这种情况会有一种弊端
// 就是main不会等待goroutine任务执行完
// 一旦main函数里的其他非goroutine的内容执行完
// 那也相当于这个程序执行完毕并且终止运行了

// 但是大部分情况下
// 我们要并发的执行任务 它们需要的时间是没办法预测的
// 这时候就需要用到通道了(channel)
// 通道可以给我们提供安全快速的执行goroutine的操作
// 通道的创建: 利用内置的make函数
// 语法: 通道名 := make(chan 类型)
// c := make(i int)

// 通道的发送和接受
// 使用左箭头操作符 <- 向通道发送值 或 从通道接收值
// - 向通道发送值 c <- 99
// - 向通道接受值 i <- c

// 发送操作会等待直到另一个goroutine尝试对该通道进行接收操作为止
// - 执行发送操作的goroutine在等待期间将无法执行其他的操作
// - 未在等待通道操作的goroutine仍然可以继续自由的运行

// 执行接受操作的goroutine将等待直到另一个goroutine尝试向该通道
// 进行发送操作为止

// 例子
func main() {
	// 创建通道
	c := make(chan int)

	for i := 0; i < 5; i++ {
		// 创建了5个goroutine任务
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		// 向通道接受值
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}

	main2()
}

// 该goroutine如果接受一个通道作为参数
// 那么声明语法为 通道名 chan 通道类型
func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Second * 3)
	fmt.Println("...", id, "snore ...")
	// 向通道发送id
	c <- id
}

// 使用select处理多个通道
func main2() {
	// 等待不同类型的值
	// time.After函数它会返回一个通道 该通道在指定时间后会接收到一个值
	// (发送该值的goroutine是Golang运行时的一部分)

	// select 和 switch 有点像
	// - 该语句包含的每个case都持有一个通道 同来发送或接受数据
	// - select会等待直到某个case分支的操作就绪 然后就会执行该case分支
	// - select语句在不包含任何case的情况下将永远等下去

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go newSleepyGopher(i, c)
	}

	// 超时时间为两秒
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		// 如果两秒内完成的通道则走以下分支
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		// 如果超过两秒完成的通道则停止等待并且走以下分支
		case <-timeout:
			// ! 注意 即使已经停止等待了goroutine 但只要main函数还没返回
			// ! 仍在运行的goroutine将会继续占用内存
			fmt.Println("my patience ran out")
			return
		}
	}
}

func newSleepyGopher(id int, c chan int) {
	// 阻塞0~4秒之内的任意时间
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Microsecond)
	c <- id
}

// nil通道
// 如果不适用make函数初始化通道
// 那么通道变量的值就是nil(零值)
// 对nil通道进行发送或接收不会引起panic(恐慌/奔溃)
// 但会导致永久阻塞
// 对nil通道执行close函数 则会引发panic(恐慌/奔溃)
// nil通道的用处:
// - 对于包含了select语句的循环 如果不希望每次循环都等待
//	 select所涉及的所有通道 那么可以先将某些通道设置为nil
//   等到发送值准备就绪之后 再将通道变成一个非nil值并执行发送操作
