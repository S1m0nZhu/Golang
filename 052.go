/*
可变参数函数
调用参数时，参数的个数可以是任意个
可变参数必须在参数列表的最后位置，在参数名和类型之间三个点表示可变参数函数
 */
package main

import "fmt"

func demo(name string,hover ... string){
	fmt.Println("woshi",name,"like")
	for i,n := range hover {
		fmt.Println(i,n)
	}
}

func main() {
	demo("cxk","唱","跳","rap","篮球")
}
