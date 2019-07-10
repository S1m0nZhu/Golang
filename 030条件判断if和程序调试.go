/**
* @Author: S1m0n
* @Date: 19-7-4 下午6:56
 */
package main

import "fmt"

func main() {
	score := 65
	//多if 无论前面的是否满足，所有的if都判断
	if score >= 60 {
		fmt.Println("及格")
	}
	if score < 60{
		fmt.Println("不及格")
	}
}
