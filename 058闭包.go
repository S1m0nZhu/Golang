package main

import "fmt"
//在另一个函数中访问其他函数的局部变量
func main() {
	//f为返回值函数
	f:=closure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	f1:=closure()
	fmt.Println(f1())
	fmt.Println(f1())
	//f2:=closure()
	fmt.Println(f())
}

func closure() func() int {
	i:=1 //重新调用会在此处初始化i
	return func() int {
		fmt.Printf("%p\n",&i)
		i=i+1//调用f和f1时i的指向不一样
		return i
	}
}