package test

import (
	"fmt"
	"testing"
)

type IPerson interface {
	Say() string
}
type liyuanqi struct{}

//
func NewPerson() IPerson {
	return liyuanqi{}
}

// liyuanqi 这个struct 实现了IPerson接口内的所有方法，就是说它实现了IPerson接口,它是隐式实现，
// 没有像Java/PHP 要implements关键字指定实现哪个接口
func (liyuanqi) Say() string {
	fmt.Println(" say ", "liyuanqi")
	return "liyuanqi"
}

type arr struct {
	Id        int
	OrderCode string  `json:"order_code"`
	Mcode     string  `json:"mcode"`
	Type      int     `json:"type"`
	Radio     float64 `json:"radio"`
	Mycat     int     `json:"mycat"`
}
type arr2 struct {
	Mcode string  `json:"mcode"`
	Radio float64 `json:"radio"`
}

func Test_Interface(t *testing.T) {
	var ex []arr
	ex = []arr{
		{1, "ORD00000000001", "m1", 1, 30.0, 125}, {2, "ORD00000000001", "m2", 2, 67.0, 125},
		{3, "ORD00000000001", "m3", 2, 3.0, 216},
		{4, "ORD00000000002", "m3", 2, 3.0, 216},
	}

	m1 := make(map[string][]arr)
	for _, a := range ex {

		m1[a.OrderCode] = append(m1[a.OrderCode], a)
	}
	var radio float64
	for _, a := range m1["ORD00000000001"] {

		if a.Type == 2 && a.Mycat == 216 {
			radio += a.Radio

		}

	}
	for i, a := range m1["ORD00000000001"] {

		if a.Type == 1 {
			m1["ORD00000000001"][i].Radio = a.Radio + radio
			t.Log(a.Radio, radio, m1["ORD00000000001"][i].Radio)
		}

	}
	t.Log(radio)

	t.Log(m1)

}

type xushuhui struct{}

//
func NewPerson2() IPerson {
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

}

func isEqual() {
	a := new(struct{})
	b := new(struct{})

	fmt.Printf("a %p b %p", a, b)
	fmt.Println()
	fmt.Printf("a %p b %p", &a, &b)
	print(" print ", a == b)
	println(" println", a == b)
	fmt.Println(" fmt ", a == b)
	print(" print 2 ", &a == &b)
	println(" println2", &a == &b)
	fmt.Println(" fmt2 ", &a == &b)
	return
}
