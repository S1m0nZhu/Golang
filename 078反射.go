/**
* @Author: S1m0n
* @Date: 19-7-13 上午10:10
反射
reflect包提供了运行时反射，程序运用过程中动态操作结构体
当变量存储结构体属性名称，想要对结构体的这个属性赋值或查看时，就可以使用反射
反射还可以用作判断变量类型
整个reflect包中最重要的2个类型reflect.Type类型 reflect.Value值

获取Type和Value的函数
reflect.TypeOf(interface{})返回Type
reflect.ValueOf(interface{})返回Value

 */
package main

import (
	"fmt"
	"reflect"
)

type People struct{
	Name string `xml:"name"`//标记通过反射
	Address string
}
func main() {
	//a:="name"
	//peo:=People{"simon"}
	//peo.name
	//
	//a:=1.5
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.ValueOf(a))
	//
	////获取结构体属性的值
	//peo:=People{"S1m0nZhu","上海"}
	//v:=reflect.ValueOf(peo)
	//fmt.Println(v.NumField())
	//fmt.Println(v.FieldByIndex([]int{0}))
	//
	//content:="name"
	//fmt.Println(v.FieldByName(content))

	////设置属性的值
	//content := "Name"
	//peo:=new(People)//new指针类型
	//v:=reflect.ValueOf(peo).Elem() //传参是值类型，用指针，形参在改变，实参也跟随改变,Elem（）获取指针指向地址的封装
	//
	////需要修改属性的内容时，要求结构体中属性名首字母要大写才可以设置
	//fmt.Println(v.FieldByName(content).CanSet())
	//v.FieldByName(content).SetString("zhumeng")
	//v.FieldByName("Address").SetString("Shanghai")
	//fmt.Println(peo)

	//通过反射获取标记
	t:=reflect.TypeOf(People{})
	fmt.Println(t.FieldByName("Name"))
	name,_:=t.FieldByName("Name")
	fmt.Println(name.Tag)
	fmt.Println(name.Tag.Get("xml"))
}

