package main

import "fmt"

func main() {
	// const (
	// 	a = iota  //自动从0开始递增
	// 	b
	// 	c
	// 	d
	// )

	// const (
	// 	a = iota  //自动从0开始递增
	// 	b = iota
	// 	c = iota
	// 	d = iota
	// )

	// const (
	// 	a = iota  //自动从0开始递增
	// 	b
	// 	c = 3
	// 	d     // 不写 默认是上面的值
	// )

	// const (
	// 	a = iota //自动从0开始递增
	// 	b
	// 	c = 3
	// 	d        // 不写 默认是上面的值 3
	// 	e = iota // 打印的是行的号 4
	// )

	const (
		a = 1 //自动从0开始递增
		b
		c 
		d         
		e    //打印都是1

	)

	fmt.Printf("%d %d %d %d %d", a, b, c, d, e)
}
