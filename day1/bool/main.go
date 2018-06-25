package main

import "fmt"

func main(){
	var b = false
	var (
		c = true
		d bool = false
	)

	_ = d  // d 赋值给下划线
	if !b {
		fmt.Print("b is false\n")
		fmt.Printf("格式化打印 %t\n",b)
	}

	if !b && c {
		fmt.Print("result both true\n")
	}

	if b || c{
		fmt.Print("b or c is true\n")
	}
}