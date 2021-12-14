package test

import (
	"fmt"
	"testing"
)

type A struct {
	Field1 string
}
type B struct {
	F int
}

func Test_a_test(t *testing.T) {
	// a := new(struct{})
	// b := new(struct{})

	// fmt.Println(a == b)

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