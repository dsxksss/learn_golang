package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 优化前
	// s, sep := "", ""

	// for _, v := range os.Args[1:] {
	// 	s += sep + v
	// 	sep = " "
	// }

	// fmt.Println(s)

	fmt.Println("What's your name?")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Printf("hello %v this is your args\n", strings.TrimSpace(name))

	// 优化后
	fmt.Println(strings.Join(os.Args[1:], " "))

	// 运行 : go build cliTool/main.go && ./main abc 123 xx=a
}
