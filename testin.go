package main

import (
	"fmt"
	"reflect"
)

func main() { testRef() }

func testRef()  {
	var x float32 = 11.11
	//将 x 变量设置为反射对象 这里用到了指针
	p := reflect.ValueOf(&x)
	//获取标签字段 也就是float32类型
	v := p.Elem()
	//修改值
	v.SetFloat(13.14)
	fmt.Println("reflect 修改后的值: ", v)

}

//定义空接口
func testNullInte()  {
	// 定义a为空接口
	var a interface{ }
	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	fmt.Println("i赋值给a接口 ：",a)
	a = s
	fmt.Println("s赋值给a接口 ：",a)

}



func testEating() {
	var man EatFood

	man = new(WaiguorenFood)
	fmt.Println( man.foodName())
	fmt.Println( man.tableware())

	man = new(ChinaFood)
	fmt.Println( man.foodName())
	fmt.Println( man.tableware())
}
//定义一个食材的接口
type EatFood interface {
	//食材名称
	foodName() string
	//食材餐具
	tableware() string
}
//定义一个歪果仁食材结构体
type WaiguorenFood struct { }

func (waiguoren WaiguorenFood) foodName() string {
	return "牛肉"
}
func (waiguoren WaiguorenFood) tableware() string {
	return "刀叉"
}

//定义一个国人结构体
type ChinaFood struct { }

func (china ChinaFood) foodName() string {
	return "猪肉"
}
func (china ChinaFood) tableware() string {
	return "筷子"
}
