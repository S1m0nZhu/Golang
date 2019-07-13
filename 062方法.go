/**
* @Author: S1m0n
* @Date: 19-7-4 下午2:49
方法
方法和函数语法比较像，函数属于包，通过包调用函数
方法属于结构体，通过结构体变量调用
默认是函数，隶属于包，所以需要添加标识，告诉编译器属于哪个结构体

调用方法时就把调用者赋值给接受者
 */
package main

import "fmt"

/*
type People struct{
	Name string
	Weight int
}

func(p *People) run(){ //方法的接收者；加*把peo的地址赋给p
	fmt.Println(p.Name,"正在跑步，当前体重为",p.Weight)
	p.Weight-=1
}

func main(){
	peo:= People{"zhangsan",100}//peo传递给p,People不带*，传递的是值副本
	peo.run() //方法的调用者需要跟随变化，指针，不跟随变化，值类型
	fmt.Println("跑步后的体重",peo.Weight)
}
//调用者