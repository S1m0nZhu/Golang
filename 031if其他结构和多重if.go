/**
* @Author: S1m0n
* @Date: 19-7-4 下午7:21
if……else结构
满足互斥条件时使用

if多重嵌套

if……else if…… else

 */

package main

import "fmt"

func main() {
	score := 73
	if score >= 60{
		if score >= 90{
			fmt.Println("优秀")
	}
		if score >= 80 && score <= 90{
			fmt.Println("良好")
		}
		if score >= 70 && score <= 80{
			fmt.Println("中等")
		}
		if score >= 60 && score <= 70{
			fmt.Println("及格")
		}
	} else {
		fmt.Println("不及格")
	}
	
}
