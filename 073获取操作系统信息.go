/**
* @Author: S1m0n
* @Date: 19-7-10 下午5:50
标准库提供不依赖平台的操作系统接口
设计为Unix风格，错误处理时Go风格，失败调用会返回错误值而非错误码，通常错误值里包含更多信息
os包及其子包功能
os/exec包，负责执行外部命令
os/signal对输出信息的访问
os/user通过名称或ID查询用户账户

Uid 用户id
Gid 所属组id
Username 用户名
Name 所属组名
HomeDir 用户对应文件夹路径
 */
package main

import (
	"fmt"
	"os/user"
)

func main() {
	//a,error:= user.Current()
	//a,error:= user.Lookup("pyvip")
	a,error:= user.LookupId("1000")

	if error != nil{
		fmt.Println(error)
		return
	}
	fmt.Println(a.Uid)
	fmt.Println(a.Name)
	fmt.Println(a.Gid)
	fmt.Println(a.HomeDir)
	fmt.Println(a.Username)
}
