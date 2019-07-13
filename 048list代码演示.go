/**
* @Author: S1m0n
* @Date: 19-7-8 下午10:25
操作List
直接使用container/list包下的New（）新建一个空的List
Go语言标准库中提供了很多双向链表中添加元素的函数
取出链表中元素
移动元素的顺序
删除元素


 */
package main

import (
	"container/list"
	"fmt"
)

func main() {
	mylist:=list.New()
	fmt.Println(mylist)

	//往头尾加
	mylist.PushFront("a")
	mylist.PushBack("b")
	mylist.PushBack("c")
	mylist.InsertBefore("d",mylist.Back())
	mylist.InsertAfter("e",mylist.Front()) //向链表后插入

	//移动元素的顺序
	mylist.MoveBefore(mylist.Front(),mylist.Back())


	//遍历
	for e:=mylist.Front();e!=nil;e=e.Next() {
		fmt.Println(e.Value,"")
	}

	n:=3
	first:=mylist.Front()
	for i:=0;i<n;i++{
		first=first.Next()
	}
	fmt.Println(first.Value)


	}

