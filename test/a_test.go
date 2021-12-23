package test

import (
	"fmt"
	"testing"
)
type IPerson interface {
	Say()string
}
type liyuanqi struct{}
// 
func NewPerson()IPerson{
	return liyuanqi{}
}
// liyuanqi 这个struct 实现了IPerson接口内的所有方法，就是说它实现了IPerson接口,它是隐式实现，
// 没有像Java/PHP 要implements关键字指定实现哪个接口
func (liyuanqi) Say() string {
	fmt.Println(" say ", "liyuanqi")
	return "liyuanqi"
}

func Test_Interface(t *testing.T) {
	// 作为调用方，我直接调用NewPerson，获取接口，然后调用接口的say方法,我不管里面是什么，我只管调用，换句话说，我还可以写个xushuhui struct
	// 同样实现say方法，然后调用NewPerson2，同样执行say 方法
	person := NewPerson()
	person.Say()
	person2 := NewPerson2()
	person2.Say()

}
type xushuhui struct{}
// 
func NewPerson2()IPerson{
	return xushuhui{}
}
// liyuanqi 这个struct 实现了IPerson接口内的所有方法，就是说它实现了IPerson接口,它是隐式实现，
// 没有像Java/PHP 要implements关键字指定实现哪个接口
func (xushuhui) Say() string {
	fmt.Println(" say ", "xushuhui")
	return "xushuhui"
}
type A struct {
	Field1 string
}
type B struct {
	F int
}

func Test_a_test(t *testing.T) {


	isEqual()


}
func isEqual(){
	a := new(struct{})
	b := new(struct{})
	
	
	fmt.Printf("a %p b %p",a,b)
	fmt.Println()
	fmt.Printf("a %p b %p",&a,&b)
	print(" print ", a== b)
	println(" println",a==b)
	fmt.Println(" fmt ", a== b)
	print(" print 2 ", &a== &b)
	println(" println2",&a== &b)
	fmt.Println(" fmt2 ",&a== &b)
	return 
}