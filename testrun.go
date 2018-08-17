package main
import (
	"runtime"
	"fmt"
	"time"
)
func main() {
	c1 := make(chan string, 1)
	testTimeout(c1)

}
func testTimeout( s chan string){
	go func() {
		time.Sleep(time.Second * 2)
		s <- "测试超时数据"
	}()
	select {
	case res := <-s:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("哇塞 超时了")
	}
}


func testSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
			fmt.Println("x == ",x)
			fmt.Println("y+x == ",y+x)
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}


func initChannRang(n int,c chan int)  {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	//关闭协程
	close(c)
}







func initChann()  {
	//创建一个channel，数据类型为int 缓存容量为2
	des:= make(chan int,1)
	//发送流
	des <- 1
	//发送流 数据为2
	des <- 2

	fmt.Println("des 地址值:", des)
	fmt.Println("des:",<- des)
	fmt.Println("des:",<- des)
	}
func testRoutine(s string)  {
	for i := 0; i < 5; i++ {
		//runtime.Gosched()这个函数主要是用于让出CPU时间片。
		// 假设A碰到代码runtime.Gosched()就把任务交给B，A间歇，B继续运行。
		runtime.Gosched()
		fmt.Println(s)
	}
}