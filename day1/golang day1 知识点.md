# golang学习

## 环境搭建

Go是的项目统一用一个文件夹进行性管理，然后在系统的环境变量中将该文件的路径加到GOPATH
自己的文件在src中,在src中进行统一的管理

## 第一个程序

## package  main 生成可执行程序

```go
package main

import(
    "fmt"
)

//main函数是入口函数
//左括号 写在func 的同一行
func main(){
    fmt.Print("hello world")    //没有换行
    fmt.Println("hello world ") //Println自动换行
    fmt.Printf("%d",100)        //Printf 格式化
}
```

注意其中的Print 首字母大写

## 不同文件中函数调用

只要在同一个package中 如在hello.go 调用world.go中的test函数，在同一个包 package main下面能够进行调用

注意打包的时候要对目录进行打包，单独对hello.go进行打包会报错 undefined test
执行 go build .\hello\ 生成hello.exe 执行hello.exe 会打印test

## go命令

**go build** 编译生成二进制文件，go build day1/hello 编译指定的目录

**go install** 编译并安装,安装可执行文件到bin目录 go install day1/hello 指定目录进行编译安装，不指定目录会将src下面的全部编译安装，在写的时候不用以src开头,执行命令后生成的文件存放到bin文件夹中

**gofmt -w .** 格式化代码

## 变量

### 单个定义

var a int
var b string
var c bool
var d int = 1

### 批量定义

var (
   a int
   b string
   c bool
   d int = 1
)

## 常量const 只读不能修改

使用const定义常量定义自增
const(
    a = 1
    b = 2
    c = 3
)

专业写法
const(
    a = iota  //iota 是与行数有关
    b
    c
)

const(
    a = 1
    b
    c    // 后面的都是1
)

## 数据类型

### 布尔类型

a. var b bool 默认为false和 var b bool = true 和 var b = false  
b. 操作符 == 和 !=
c. 取反操作符： !b
数据类型和操作符
d. && 和 || 操作符
e. 格式化输出占位符: %t

定义了暂时不同的变量 赋值给下划线 就不会报错，导入包的时候同样如此

import (
    _ "fmt"
)

## 整数

根据数的范围进行选择
int8  一个字节 256  有符号 -127~128
int16
int32
int64

uint8
unit16
uint32
uint64

int uint  与操作系统平台相关

## 浮点数

float32
float64 4个字节

### 重点

Go是强类型的，不同类型的不能运算
需要强制类型转换

%d 整数占位符
%s 字符串占位符
%x 十六进制占位符
%v 万能占位符
%c 字符  单引号 只能是一个字符

## 字符串

使用双引号
使用反引号，原样输出
单引号 %c  单个字符

字符和字符串的区别

var a bytes
var c rune
c = '中'  单引号  单个字符

底层存储的是编码

## 程序结构

访问权限 内部小写可以访问，外部不能访问，使用大写外部能访问 
自己写的calc 调用的时候用calc.Add()

## golang 语言特性

### 垃圾回收

### 天然并发

golang的线程是在用户态实现的轻量级的线程，类似协成

go关键字 
go calc.Add(2,3)

CSP 多线程+管道  

Java等是操作系统级别的线程
:= 初始化并赋值

channel 线程之间通信 FIFO
死锁检测(新版本)

### 多返回值

func test() (int,int){
    return a,b
}

## 包

包的概念,在同一个包下面写的go程序，可以在外部进行调用，导入的时候需要导入路径，能提高代码的复用性

### 静态库

自己写的calc包执行 go build 然后执行go install 会在pkg中生成calc.a的静态库，相当于缓存