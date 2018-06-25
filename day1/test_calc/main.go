package main

import (
	"fmt"
	"day1/calc"
)

func main(){
	var res int
	var res1 string
	res = calc.Add(2,3)  //调用静态库中的函数 
	fmt.Print(res)
	res1 = calc.Test1() //有返回值 需要接收
	fmt.Print(res1)
}