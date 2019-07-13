/*
匿名函数
正常函数通过名称多次调用
，匿名函数由于没有函数名，所以大部分情况都是在当前位置声明并立即调用（函数变量除外）
匿名函数声明完需要调用，在函数结束大括号后面紧跟小括号
匿名函数声明在其他函数内部*/
package main

import (
	"fmt"
)

func main() {
	func(){
		fmt.Println("无参无返回值匿名函数")
	}()

//有参数无返回值的匿名函数
	func(name string){
	fmt.Println("名字为",name)
	}  ("S1m0n")

//无参有返回值的匿名函数
	name:=func() string{
		fmt.Println("有返回值的匿名函数")
		return "zhumeng"
	}()//此处有调用
	fmt.Println(name)
}

