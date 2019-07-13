/*
上节中，每次创建指针必须再额外创建一个变量，操作比较麻烦
new函数
直接创建一个类型指针

当我们想直接创建一个指针变量，且在直接使用，而不想通过别的变量来赋值
 */
package main

import "fmt"

func main() {
	a:=new(int)//直接使用
	fmt.Println(a)

	*a = 123
	fmt.Println(*a)

	var b *int //空指针
	fmt.Println(b)
}
