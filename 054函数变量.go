/*
函数变量
go语言中，函数也是一种类型，函数有多少种形式，函数变量就有多少种写法
定义完函数变量，可以使用匿名函数赋值，也可以使用已经定义好的函数赋值
函数定义后与普通函数调用语法相同，变量名就是普通函数声明的函数名

函数类型变量是除了slice/map/channel/interface/外第五种引用类型

变量可以当作函数的返回值，而函数变量也可以作为函数的参数或返回值
函数作为参数时，类型写成对应的类型即可
函数作为返回值
 */
package main

import "fmt"

/*func main() {
	c()
	fmt.Println(c)

	/*var b func()
	b = func() {
		fmt.Println("b")
	}
	b()
	b=c //普通函数声明后的函数名
	b()
	*/
/*	var x func()
	x=c
	fmt.Printf("%p %p",c,x)
}
func c(){
	fmt.Println("a")
}*/


//函数作为参数
/*func main(){
	/*r:=ring.New()
	r.Do(func(i interface[]) {

	}

	mydo(func(name string){//形参
		fmt.Println(name)
	})
}

func mydo(arg func(name string)){
	fmt.Println("执行,mydo")
	arg("goland")//实参
}*/


//函数作为返回值
func main(){
	result:=a() //函数传递给变量
	r2:=result()//想要使用函数就调用，调用结果再传递给r2
	fmt.Println(r2)
}

func a() func() string {
	return func() string {
		return "man"
	}
}