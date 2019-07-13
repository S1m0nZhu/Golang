package main

import "fmt"

func main() {
	a:=119
	b:="zhumeng" //函数参数 与名字没关系
	c:=[]int{1,2} //c传递的是地址，arr👈指向c
	d:=119
	demo(a,b,c,&d)//值类型传递的是副本，形参改变不影响实参
	fmt.Println(a,b,c,d)
}


func demo(i int ,s string, arr []int, content *int) {
	i = 110
	s = "simon"//值传递时是深拷贝，改的是副本，形参改变不影响实参
	arr[0] = 3
	arr[1] = 5//引用类型，改的是指向的同一块内存地址
	*content =110 //传递指针就是传递地址；如果希望值类型数据在修改形参时实参跟随变化，可以把参数设置为指针类型
}