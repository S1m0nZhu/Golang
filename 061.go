/**
* @Author: S1m0n
* @Date: 19-7-4 上午11:18
结构体指针
由于结构体是值类型，在方法传递时希望传递结构体地址，可以使用结构体指针完成
可以结合new(T)函数创建结构体指针
*/
/*package main

import "fmt"

type People struct{ //结构体变量是指针类型
	Name string
	Age int

}*/

//func main() {
	//peo:=new(People)
	//fmt.Println(peo)
	//fmt.Println(peo==nil)
	//
	//peo.Name="s1mm00000000n"
	//peo.Age=17
	//fmt.Println(peo)
	//
	//peo2 := peo
	//fmt.Println(peo2)
	//
	//peo2.Name="zhuyang"
	//fmt.Println(*peo,*peo2)

	//peo := &People{"zhuyangggg",15}
	//peo2 := &People{"zhuyangggg",15}//创建指针类型方式2；属性相同，地址不同
	//fmt.Printf("%p%p\n",peo,peo2)
	//fmt.Println(peo==peo2)//双等比较，值类型比较的是值；指针类型比较的是地址
}


package main

import "fmt"

type People struct{
	Name string
	Age int
}

func main() {
	//peo := new(people)//单纯的用结构体指针，new方法
	//fmt.Println(peo)
	//fmt.Println(peo==nil)
	//
	//peo.Name = "simon"
	//peo.Age = 17
	//fmt.Println(peo)
	//
	//peo2:=peo
	//fmt.Println(peo2)
	//
	//peo2.Name = "zhumeng"
	//fmt.Println(peo2,peo)

	//直接创建指针类型
	peo := &People{"zhumeng",17}
	peo3 := &People{"zhumeng",17}
	fmt.Printf("%p %p\n",peo,peo3)
	fmt.Println(peo==peo3)
	//
	peo6:= People{"zhumeng",17}
	fmt.Println(*peo==peo6)//*取地址中的值

}
