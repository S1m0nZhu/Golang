/**
* @Author: S1m0n
* @Date: 19-7-10 下午7:54
一、os包内容介绍
权限控制 可读可写
Filemode表示文件权限，本质时uint32,可取值都以常量形式提供
Fileinfo是一个interface表示文件信息

二、资源路径问题
在获取资源路径时分为相对路径和绝对路径
相对路径在Go语言中时GOPATH，也就是项目的根目录
绝对路径：磁盘根目录开始表示资源详细路径的描述

标准库提供2钟创建文件夹的方式
os.mkdir
os.mkdirall
os.Create 创建文件时要求文件所在的目录必须已经存在，如果文件已存在则会创建一个空文件覆盖之前的文件

 */
package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件夹
	//err := os.MkdirAll("/home/pyvip/ab", os.ModeDir)
	//if err != nil {
	//	fmt.Println("文件夹创建失败")
	//	return
	//}
	//fmt.Println("文件夹创建成功")

	//创建文件
	f,err := os.Create("/home/pyvip/b.txt")
	if err!=nil{
		fmt.Println("create failed")
		return
	}
	fmt.Println("create success",f)

	//重命名文件
	//err:= os.Rename("/home/pyvip/ab","/home/pyvip/bc")
	//if err != nil {
	//	fmt.Println("重命名失败")
	//	return
	//}
	//fmt.Println("文件夹重命名成功")

	//重命名文件
	//err := os.Rename("/home/pyvip/b.txt", "/home/pyvip/bbc.txt")
	//if err != nil {
	//	fmt.Println("failed")
	//	return
	//}
	//fmt.Println("success")

	//打开文件
	//f,err := os.Open("/home/pyvip/bbc.txt")
	//if err != nil {
	//	fmt.Println("文件获取失败", err)
	//	return
	//}
	//fileInfo, err := f.Stat()
	//if err!=nil{
	//	fmt.Println("文件获取失败",err)
	//	return
	//}
	//fmt.Println(fileInfo.Size())
	//fmt.Println(fileInfo.ModTime())
	//fmt.Println(fileInfo.Mode())
	//fmt.Println(fileInfo.IsDir())
	//fmt.Println(fileInfo.Name())

	//删除文件、目录需为空
	err1:=os.RemoveAll("/home/pyvip/abc")
	if err1 != nil{
		fmt.Println("删除失败",err1)
		return
	}
	fmt.Println("删除成功")

	//删除文件、目录
}


