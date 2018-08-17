package main
import "fmt"

func main() { testEat() }
func testEat() {
	var eat Eat
	eat = new(Waiguoren)
	eat.eating()
	eat = new(China)
	eat.eating()
}
//定义一个吃的接口
type Eat interface { eating() }
//定义一个歪果仁结构体
type Waiguoren struct { }
func (waiguoren Waiguoren) eating() { fmt.Println("歪果仁用刀叉")}
//定义一个国人结构体
type China struct { }
func (china China) eating() { fmt.Println("中国人用筷子")}
