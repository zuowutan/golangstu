package main

import "fmt"

func main()  {
	var c1, c2  int = 30,40
	var i1, i2 int = 20,30

	select {
	case i1 > c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 > i2:
		fmt.Printf("sent ", i2, " to c2\n")

	default:
		fmt.Printf("no communication\n")
	}
/*	var Xiaowu Student
	var Laowang Student
	var Lao Student

	Lao.name = "大梦谁先觉"

	Xiaowu.number = 18218013313
	Xiaowu.age = 26
	Xiaowu.name = "小武"

	Laowang.age = 27
	Laowang.name = "老王"
	Laowang.number = 435725233

	//useObject(Lao)
	//useObject(Laowang)

	//pointAddre(Lao)
	//pointAddre(Laowang)

	pointAddre(Xiaowu)*/

}

func pointAddre(student Student){

	var struct_pointer *Student
	struct_pointer = &student

	var dataNumber = struct_pointer.number
	var dataAge = struct_pointer.age
	var dataName = struct_pointer.name

	println(dataNumber)
	println(dataAge)
	println(dataName)

	println("分割线")

/*	println("对象的指针是：", &student)
	println("对象的地址值是：", &student)*/
}
//使用对象
func useObject(student Student)  {
	println(student.name)
	println(student.number)
	println(student.age)
}

type Student struct {
	age int
	name string
	number int
}
