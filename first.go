 package main

import "fmt"

const POINT_SIZE = 10

var data = "嘿嘿嘿"

//初始化切片 一
var sliceName = make([]int, 8)

func main() { changeStyle() }

func changeStyle()  {
	var a int = 17
	var b int = 5
	var data float32
	//将int类型的转换为float类型
	data = float32(a)/float32(b)
	fmt.Printf("data 的值为: %f\n",data)
}

func initMap() {
	//	创建Map 方式一：
	var mMap map[string]string = make(map[string]string)
	//向Map中添加元素
	mMap [ "1" ] = "巴西"
	mMap [ "2" ] = "阿根廷"
	mMap [ "3" ] = "意大利"
	mMap [ "4" ] = "德国"
	mMap [ "5" ] = "日本"

	/* 遍历Map */
	for mKey := range mMap {
		fmt.Println(mKey, " 索引对应的结果：", mMap [ mKey ])
	}




	//	创建Map 方式二：
	rating := map[string]float32{"Go":18, "Java":18, "Android":18, "C++":16 }
	fmt.Println(len(rating))


}
func addAndCopy() {
	//初始化切片
	var data []int
	//切片添加单个元素
	data = append(data, 0)
	data = append(data, 1)
	//切片添加多个元素
	data = append(data, 2, 3, 4)
	var i int
	for i = 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
	fmt.Println("----------")
	for _, i := range data {
		fmt.Println(data[i])
	}
}

func initSlice() {
	//初始化切片 二
	s := []int{1, 2, 3, 4, 5, 6}
	//获取切片的长度
	fmt.Println(len(s))
	fmt.Println("-------")
	//获取切片的容量
	fmt.Println(cap(s))

}

func inputContent() {
	var user User
	var data *User

	user.account = "182180133"
	user.password = "123456"
	user.secondPassword = "123456"

	data = &user

	fmt.Println(data)
	fmt.Println(data.account)
	fmt.Println(data.password)
	fmt.Println(data.secondPassword)

	//
}

//显示用户信息
func showUserData(user *User) {

	fmt.Printf("User account : %s\n", user.account)
	fmt.Printf("User password : %s\n", user.password)
	fmt.Printf("User secondPassword : %s\n", user.secondPassword)
}

func registApi(user User) {
	if len(user.account) > 0 && len(user.password) > 0 && len(user.secondPassword) > 0 {
		if user.password == user.secondPassword {
			//这里要做判断 如果之前已经注册的要提示用户已经注册 没有注册的则进行后续的判断
			fmt.Println("用户注册成功")

		} else {
			fmt.Println("两次密码输入不一致！")
		}
	} else {
		fmt.Println("账号/密码输入为空！")
	}

	showUserData(&user)

}

//定义一个Go结构体
type GoBean struct {
	//这是结构体具体的字段
	name string
	old  int
	num  string
	love string
}

//

//入：
//

type User struct {
	//用户名
	account string
	//密码
	password string
	//再次输入密码
	secondPassword string
}

func useGoBean() {
	var xiaozhuliuxing GoBean

	xiaozhuliuxing.name = "小猪流星"
	xiaozhuliuxing.old = 25
	xiaozhuliuxing.num = "18218013333"
	xiaozhuliuxing.love = "啪啪啪"

	fmt.Println(xiaozhuliuxing.name)
	fmt.Println(xiaozhuliuxing.old)
	fmt.Println(xiaozhuliuxing.num)
	fmt.Println(xiaozhuliuxing.love)
}

//二级指针
func ppoint() {

	var data = 100
	//一级指针
	var pointData *int
	//二级指针
	var ppointData **int

	// 指针 pointData 地址
	pointData = &data

	// 指向  指针pointData 地址
	ppointData = &pointData

	fmt.Printf("变量 data = %d\n", data)
	fmt.Printf("一级指针变量地址值 pointData = %d\n", pointData)
	fmt.Printf("指向指针的指针变量（二级指针）地址值 ppointData = %d\n", ppointData)

	fmt.Printf("一级指针的地址 对应的属性值 *ppointData = %d\n", *pointData)

	fmt.Printf("指向指针的指针变量（二级指针）地址 对应的地址值 *ppointData = %d\n", *ppointData)
	fmt.Printf("指向指针的指针变量（二级指针）地址 对应的属性值 **ppointData = %d\n", **ppointData)

}

func pointArray() {

	arr := []int{10, 100, 1000, 10000}

	var i int

	const MAX = 4

	var pointArr [MAX]*int

	for i = 0; i < MAX; i++ {
		// 将数组内属性值的地址 赋值给 指针数组  &arr[i]
		pointArr[i] = &arr[i]
	}
	for i = 0; i < MAX; i++ {
		//现在指针数组里面已经有指针（地址值）了 那么通过*指针 即可获取具体的变量
		fmt.Printf("arr[%d] = %d\n", i, *pointArr[i])
	}

}

func useData() {
	var data = "羞羞羞"
	fmt.Println("这是方法里面的全局变量,它的名字是", data)
}

func closePack() func() int {
	i := 1
	var a = 10

	println(a)
	return func() int {
		i *= 2
		return i
	}
}

func addClose(x, y int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x + y
	}
}

func closePac(a int, str string) func() string {
	return func() string {
		return str
	}
}

/* 定义交换值函数 */
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

//允许返回多个值
func returnNum(a int, s string, d float32) (int, string, float32) {
	return a, s, d
}

func addNum(a, b int) int {
	return a + b
}

// n*n 乘法表
func numTake(n int) {
	//
	var i, j int

	for i = 0; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Println(j, "*", i, "= =", i*j)
		}
		fmt.Println()
	}
}

/*	var i int
	var arrayName = [6]int{62,55,264,45451,121,892}
	for i = 0;i< len(arrayName);i++ {

		if i == 3{

			continue
		}
		fmt.Println(arrayName[i])
	}
*/

/*	for i = 1;i <= 9 ;i++ {
		for j = 1;j<=i ; j++ {
			fmt.Println(j,"*",i,"==",i*j)
		}
		fmt.Println()
	}*/

//
//
//

/*	const (
		a = iota
		b
		c = false
		d
		e = "30"
		f
		g = iota
		h
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)*/

/*	//显示定义数据类型
	const a int = 9
	//不显示定义数据类型
	const name  = "nameData"

	fmt.Println(a)
	fmt.Println(name)*/

/*	var pointName *int

	if (pointName == nil) {
		fmt.Println("这里为空指针")

	}else {
		fmt.Println("不为空")
	}
	fmt.Println(pointName)
*/
/*	var pointName *int

	pointName = &data

	fmt.Println("data的属性值：",data)

	fmt.Println("data的属性值对应的内存地址：",memoryAddress)
	fmt.Println("pointName 指针变量:属性值：",pointName)

	fmt.Println("pointName 指针变量:赋值之后的属性值：",*pointName)*/

/*	*/

func test() {

	var p int = 20 //声明实际变量
	var ip *int    //声明指针变量
	fmt.Println("变量的指针地址是:", &p)
	fmt.Println("指针变量 ip 的地址是", ip)
	//当指针变量的存储地址 实际赋值的时候
	ip = &p
	fmt.Println("指针变量 ip 赋值以后的的地址是", ip)
	fmt.Println("指针变量 ip 赋值以后的的地址，对应的变量属性值是", *ip)
	var a = 129
	var f = 123.5
	var b = false
	var s = "这是一个字符串"
	var n_i int
	var n_s string
	var n_f float64
	var n_b bool
	fmt.Println("整数型 : ", a)
	fmt.Println("浮点型 : ", f)
	fmt.Println("布尔类型 : ", b)
	fmt.Println("字符串类型 : ", s)
	fmt.Println("!--分割线--!")
	fmt.Println("默认的浮点型 : ", n_f)
	fmt.Println("默认的布尔类型 : ", n_b)
	fmt.Println("默认的字符串类型 : ", n_s)
	fmt.Println("默认的int类型 : == ", n_i)
}
