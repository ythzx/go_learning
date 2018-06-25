package main 

import "fmt"

func main(){
	var a int = 10 
	var b int32 = 10
	var c int 
	c = a + int(b)  // 强制类型转换 不同的int 类型也要进行强制转换

	fmt.Printf("%v",c) // %v 显示任意类型
	fmt.Printf("%x",c) // %x 16 进制  20 --14  
}