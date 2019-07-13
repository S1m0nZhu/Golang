/**
* @Author: S1m0n
* @Date: 19-7-11 下午9:35
Writer输出流
os.Open获取文件是以只读方式
f,err := os.OpenFile(fp, os_APPEND,0660)
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/home/pyvip/go/src/test2.txt"
	f,err := os.OpenFile(filePath,os.O_APPEND,0660)
	//defer f.Close()
	if err != nil{
		f,_=os.Create("/home/pyvip/go/src/test2.txt")
	}
	f.Write([]byte("比你优秀的人还比你努力"))
	f.WriteString("\n决心+执行力+意志力=成功")
	fmt.Println("志向与追求")
}
