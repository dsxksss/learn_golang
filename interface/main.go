package main

type log struct {
	text       string
	createTime int
}

// 实现logConfiger接口的方法
func (l log) log() string {
	return "error" + l.text
}

// 接口类型
// 为了复用 通常会把接口声明为类型
// 按约定 接口名称通常以er结尾
// 只要某个类型的方法签名和某个接口中的方法签名一致
// 那么该方法则实现了该接口

type logConfiger interface {
	log() string
	showLog()
}

// 任何满足了logConfiger这个接口的类型都可以传入进来
func shout(er logConfiger) {

}

// ! Golang中的接口的都是隐式实现的

func main() {
	logre := log{text: "message xxx", createTime: 1234560}
	logre.log()
}
