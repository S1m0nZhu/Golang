/**
* @Author: S1m0n
* @Date: 19-7-4 下午10:49
接口
接口是一组行为规范的定义
接口中只能有方法声明，方法只能有名称，参数，返回值，不能有方法体
每个接口可以有多个方法声明，结构体把接口中所有方法都重写后，结构体就属于接口类型
Go语言中接口和结构体之间的关系，是传统OOP中is-like-a的关系
定义接口的关键字时interface

接口可以继承接口，且Go语言中推荐把接口中的方法拆分成多个接口

接口中声明完方法，结构体声明完接口中的方法后，编译器认为结构体实现了接口
重写的方法要求必须与接口的方法名称、方法参数（参数名称可以不同）、返回值列表完全相同
 */

/*
package main

import (
	main2 "awesomeProject/GoPractice"
	"fmt"
)

type Eat interface {
	eat()
}

type Live interface{
	run(run int)
	Eat
	eat()
}


type Animal struct {

}

func (a *Animal) eat(){
	fmt.Println("动物在吃饭")
}


//type People struct{
	name string
}

//必须把方法都重写
func (p *main2.People) run(run int){
	fmt.Println(p.name,"正在跑步,跑了",run,"米")
}

func (p *main2.People) eat(){
	fmt.Println(p.name,"正在吃饭")
}


func main() {
	peo:= main2.People{"Simon"}
	peo.run(100)
	peo.eat()
	ani:=Animal{}
	ani.eat()
}
*/