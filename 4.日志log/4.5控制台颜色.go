package main

import "fmt"

func main() {
	// 前景色
	fmt.Println("\033[30m 黑色 \033[0m")
	fmt.Println("\033[31m 红色 \033[0m")
	fmt.Println("\033[32m 绿色 \033[0m")
	fmt.Println("\033[33m 黄色 \033[0m")
	fmt.Println("\033[34m 蓝色 \033[0m")
	fmt.Println("\033[35m 紫色 \033[0m")
	fmt.Println("\033[36m 青色 \033[0m")
	fmt.Println("\033[37m 灰色 \033[0m")
	// 背景色
	fmt.Println("\033[40m 黑色 \033[0m")
	fmt.Println("\033[41m 红色 \033[0m")
	fmt.Println("\033[42m 绿色 \033[0m")
	fmt.Println("\033[43m 黄色 \033[0m")
	fmt.Println("\033[44m 蓝色 \033[0m")
	fmt.Println("\033[45m 紫色 \033[0m")
	fmt.Println("\033[46m 青色 \033[0m")
	fmt.Println("\033[47m 灰色 \033[0m")
}
