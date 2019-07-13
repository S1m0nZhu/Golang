/**
* @Author: S1m0n
* @Date: 19-7-13 上午10:02
ioutil文件读写的工具包，通过这些函数快速实现文件的读写操作


 */
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//f,_:=os.Open("/home/pyvip/go/src/test1.txt")
	//c,_:=ioutil.ReadAll(f)
	//fmt.Println(string(c))

	//f,_:=ioutil.ReadFile("/home/pyvip/go/src/test2.txt")
	//fmt.Println(string(f))
	//
	//ioutil.WriteFile("/home/pyvip/go/src/test.txt",[]byte("这是写的数据"),0666)//先清空在写，局限性时不能尾加
	//fmt.Println("写入数据")
	//
	dir,_:=ioutil.ReadDir("/home/pyvip")
	for _,n:= range dir {
		fmt.Println(n.Name())
	}


}
