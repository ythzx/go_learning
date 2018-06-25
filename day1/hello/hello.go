package main

import (
	"fmt"
)

func main() {
	fmt.Print("hello world\n") //直接打印 没有换行
	fmt.Println("hello world") //打印并换行
	fmt.Printf("%d", 100)      //格式化打印

	var user_input int
	fmt.Print("请输入:")
	fmt.Scan(&user_input) //获取用户输入 在hello.exe 运行完后不会闪退

	test() // 调用world.go 中的test函数

	// var a int
	// var b string
	// var c bool
	// var d int = 8
	// var e string = "test"

	// var (
	// 	a int
	// 	b string
	// 	c bool
	// 	d int    = 8
	// 	e string = "test"
	// )

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
}
