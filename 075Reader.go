/*
Reader输入流
流stream，应用程序和外部资源进行数据交互的纽带
输入流、输出流都是相对于程序，外部->程序=输入，反之则反
Input Stream、Output Stream、I/O流
把流中的切片返回到切片
根据一个字符串数据产生一个字符串流
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	//r:=strings.NewReader("hello world")
	//
	//b:=make([]byte,r.Size())
	//
	//n,err:=r.Read(b)
	//if err!=nil{
	//	fmt.Println("读取数据失败，",err)
	//	return
	//}
	//fmt.Println("读取数据长度为",n)
	//fmt.Println("读取内容为",string(b))

	//从文件中读取
	f,_:=os.Open("/home/pyvip/golang/src/awesomeProject/GoPractice/test.txt")
	fileInfo,_ := f.Stat()

	b:=make([]byte,fileInfo.Size())

	f.Read(b)
	fmt.Println("文件内容是",string(b))

}

