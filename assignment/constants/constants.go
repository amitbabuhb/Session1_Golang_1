package constants

import (
	"fmt"
)

func Staticvar() {
	var a int
	a = 56
	fmt.Println(a)
	fmt.Printf("x is of type %T\n", a)

}

func Dynamicvar(){
	var b float32 = 36.6

	c := 38
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Println("Sum is ",add1(int(b),c))

}


func add1(a,b int ) (int){
	c:= a + b
	return c
}