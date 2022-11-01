package main

import "fmt"

// 方法(一种类似于函数东西)
// 在C#或JAVA中 方法属于类内容
// 但是GO中只提供了方法 而没有提供类
// GO比其他语言的方法要灵活

// 可以将方法与同包中的任何类型相关联
// 但是不能与int、float等预定义类型关联
// 方法的创建语法
// func (receiverName type) methodName() resultType {}
type newInt int

func (x newInt) newIntFunction() newInt {
	return x
}

// 上例中的newIntFunction方法虽然没有参数
// 但是它前面却有一个类型参数的接收者
// 每个方法可以有多个参数 但是只能有一个接收者
// 在方法体中 接收者的行为和其他函数一样

type kelvin float64
type celsius float64

// 调用方法
func main() {
	var k kelvin = 294.2
	var c celsius
	c = kelvinToCelsius(k)
	c = k.celsius()
	fmt.Println(c)
}

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)

}
