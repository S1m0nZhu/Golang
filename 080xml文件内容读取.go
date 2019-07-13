/**
* @Author: S1m0n
* @Date: 19-7-13 下午7:47
encoding/xml包下提供了对XML序列化和反序列化的API
使用unmarshal可以直接把XML字节切片数据转换为结构体
转换时按照特定的转换规则进行转换，且数据类型可以自动转换
通过标记实现 识别标签
如果结构体字段类型位xml.name且名为XMLNAME，unmarshal会将该元素名写入该字段
如果xml元素的属性名字匹配某个标签“.attr"为字段的字段名，或者匹配某个标签为”name,attr"的字段的标签名，unmarshal会将该属性的值写入字段
一个结构体对应一个元素切片
 */
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type People1 struct{
	XMLName xml.Name `xml:"people"`
	Id 		int  	 `xml:"id.addr"`
	Name 	string 	 `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peo := new(People1)//为了形参随实参改变，要用new传递结构体的指针
	b, err :=ioutil.ReadFile("people1.xml")
	fmt.Println(string(b))
	fmt.Println("111",err)
	err = xml.Unmarshal(b, peo)
	fmt.Println("222",err)
	fmt.Println(peo)
}
