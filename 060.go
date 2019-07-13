package main
/*结构体
将一个变量或多个变量组合在一起
是值类型
定义在函数外部，结构体名称首字母大小写影响跨包访问
如果结构体能跨包访问，属性首字母大小写影响属性是否跨包访问
 */
import (
	"fmt"
)
 type people struct {
 	Name string
 	Age int
 }

func main() {
	var peo people
	fmt.Println(peo)
	fmt.Printf("%p",&peo)

	peo = people{"simon",25}
	fmt.Println(peo)

	peo2 := people{Age:26,Name:"zhumeng"}
	fmt.Println(peo2)

	peo.Name = "mingzi"
	peo.Age = 19
	fmt.Println(peo.Name,peo.Age)

	fmt.Printf("%p %p\n",&peo,&peo2)
	fmt.Println(peo==peo2) //当成员属性完全一样时，双等比较结果一样
}