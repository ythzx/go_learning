package main

import "fmt"

//打印 字符串 二进制 十进制  十六进制 浮点数
func main(){
	fmt.Print("print string\n")
	fmt.Printf("%b\n",12)   //打印12 的二进制 1100
	fmt.Printf("%d\n",123)  //打印十进制
	fmt.Printf("%x\n",17)  //ox11 十六进制
	fmt.Printf("%f\n",1.234) //打印浮点数
}