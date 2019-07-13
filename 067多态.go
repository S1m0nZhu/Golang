/**
* @Author: S1m0n
* @Date: 19-7-4 下午11:11
多态：同一件事情由于条件不同产生的结果不同
由于Go语言中结构体不能相互转换，所以没有结构体（父子结构体）的多态，只有基于接口的多态，
这也符合Go语言对OOP的诠释

多态在代码层面最常见的一种方式是接口当作方法参数

结构体实现了接口的全部方法，就认为结构体属于接口类型，这时可以把结构体变量赋值给接口变量

重写接口时接收者为Type 和 *Type 的区别

*Type可以调用*Type和Type作为接收者的方法，所以只要接口中多个方法至少出现一个使用*Type作为接收者
进行重写的方法，就必须把结构体指针赋值给接口变量，否则编译报错

Type只能调用Type作为接收者的方法
 */
package main

import "fmt"

/*type Live interface {
	run()
	eat()
}

type People struct {
	name string
}

func (p People) run() {
	fmt.Println(p.name,"在跑步")

}

func (p People) eat() {
	fmt.Println(p.name,"在吃饭")
}
func main() {
	var live Live = People{"张三"} //直接赋值
	live.eat()
	live.run()
}
*/

type Live interface {
	run()
}

type People struct {

}

type Animal struct {

}

func (p *People) run(){
	fmt.Println("人在跑步")
}

func (a *Animal) run(){
	fmt.Println("动物在奔跑")
}

func allrun(live Live){ //allrun函数的参数是接口，接口不能实例化，通过实例化它的实现类-结构体
	live.run()
}

func main(){
	peo := &Animal{}
	allrun(peo) //给定不同的情况=参数时，代码的结果不同，这是多态的一个体现
}