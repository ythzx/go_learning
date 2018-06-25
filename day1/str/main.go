package main 

import "fmt"

func str1(){
	var b string = "hello\n\n\n"
	var c = "hello"
	fmt.Printf("b:%v and c:%s",b,c)
}

func str2(){
	var d = `测试文字 \n\n\n` // 原样输出
	fmt.Printf("%s",d)
}


func test_byte(){
	var v rune  // 
	// v = '中'    // 中文字符打印出来的十进制是20013
	v = 20013
	// fmt.Printf("v= %c",v)
	// fmt.Printf("v= %d",v)
	fmt.Printf("v= %c",v)  //输出字符
}

func main(){
	str1()
	str2()

	test_byte()
}